// amay
// Copyright (C) 2017+ skywalkerd and the project contributors
// Written by A. Jacques <andre.jacques@protonmail.com> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package fileutil

import (
  "os"

  "io/ioutil"
  "strings"
  "errors"
)

// Checks if a file exists or not
func Exists(path string) (bool, error) {
  _, err := os.Stat(path)

  if err == nil {
    return true, nil
  }

  if os.IsNotExist(err) {
    return false, nil
  }

  return true, err
}

// Checks if a specified path points to a folder or not
func IsDirectory(path string) (bool, error) {
  fileInfo, err := os.Stat(path)
  if fileInfo == nil {
    return false, err
  }
  return fileInfo.IsDir(), err
}

// Creates a file to the specified path
func CreateFile(path string) error {
  var _, err = os.Stat(path)

  if os.IsNotExist(err) {
    var file, err = os.Create(path)
    defer file.Close()
    return err
  }

  return nil
}

// Reads the content of the specified path
func ReadFile(path string) ([]byte, error) {
  exist, err := Exists(path)
  if err == nil && exist {
    file, err := ioutil.ReadFile(path)

    if err == nil {
      return file, nil
    }
  }

  return nil, errors.New("Error reading file " + path + ": file do not exist")
}

// Writes content to a specified path
func WriteFile(path string, content string) error {
  if exist, _ := Exists(path); exist {
    return ioutil.WriteFile(path, []byte(content), 0644)
  }
  return errors.New("Can't write to file: file doesn't not exists.")
}

// Creates a file to a specified destination path from a specified template path. It changes all occurrences of patterns
// by a text
func CreateFromTemplate(template string, dest string, pattern string, text string) error {
  // Creates the empty file
  err := CreateFile(dest)

  if err != nil {
    return err
  }

  // Read the content from the template
  file, err := ioutil.ReadFile(template)

  if err != nil {
    return err
  }

  fileBytes := string(file)
  if pattern != "" {
    // Change all occurrences of the pattern by the text
    fileBytes = strings.Replace(fileBytes, pattern, text, -1)
  }

  // Write the new content in the destination file
  err = ioutil.WriteFile(dest, []byte(fileBytes), 0644)

  return err
}