// amay
// Copyright (C) 2017+ skywalkerd and the project contributors
// Written by skywalkerd <andre.jacques@protonmail.com> and the project contributors
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

type IFile interface {
  Read() error
  Save() error
}

// Holds all information in regard of a file
type File struct {
  path, content string
}

// Gets the path of a file
func (f File) GetPath() string {
  return f.path
}

// Sets the path of a file
func (f File) SetPath(path string) {
  f.path = path
}

// Gets the content of a file
func (f File) GetContent() string {
  return f.content
}

// Sets the content of a file
func (f File) SetContent(content string) {
  f.path = content
}

// Reads the content of the file
func (f File) Read() error {
  bytes, err := ReadFile(f.path)

  if err == nil {
    content := string(bytes)
    f.SetContent(content)
    return nil
  }

  return err
}

// Saves the file
func (f File) Save() error {
  err := CreateFile(f.GetPath())

  if err == nil {
    err = WriteFile(f.GetPath(), f.GetContent())
  }

  return err
}