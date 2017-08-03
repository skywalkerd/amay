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

import "testing"

func TestFileExist(t *testing.T) {
  file := File { path: "test/file1.txt" }

  file.Read()

  expected := `This is a file.

This is a {pattern}.

This is not a {pattern.`

  if file.Content() != expected {
    t.Error("Can't read file content.")
  }
}

func TestFileNotExist(t *testing.T) {
  file := File { path: "test/fileNotExist.txt", content: "This is awesome\nThis is amazing" }

  if err := file.Save(); err != nil {
    t.Error("Can't save file content.")
  }

}