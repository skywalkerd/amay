package fileutil

import "testing"

// Test for the existence
func TestExistsFile1(t *testing.T) {
  exist, err := Exists("test/file1.txt")

  if !exist || err != nil {
    t.Error("This file should exists.")
  }
}

// Test for the existence
func TestExistsFolderA(t *testing.T) {
  exist, err := Exists("test/folderA")

  if !exist || err != nil {
    t.Error("This file (folder) should exists.")
  }
}

// Test for the non-existence
func TestExistsFail(t *testing.T) {
  exist, _ := Exists("bouya")

  if exist {
    t.Error("This file/folder should not exists.")
  }
}

// Test for the fact it is a directory
func TestIsDirectoryFolderA(t *testing.T) {
  isDir, err := IsDirectory("test/folderA")

  if !isDir || err != nil {
    t.Error("It should be a folder.")
  }
}

// Test for the fact it is not a directory
func TestIsDirectoryFile1(t *testing.T) {
  isDir, _ := IsDirectory("test/file1.txt")

  if isDir {
    t.Error("It is not a folder.")
  }
}

// Test for the fact it is not a directory
func TestIsDirectoryFail(t *testing.T) {
  isDir, _ := IsDirectory("bouya")

  if isDir {
    t.Error("This file/folder does not exist.")
  }
}

// Test to create a file
func TestCreateFile2(t *testing.T) {
  err := CreateFile("test/file2.txt")

  if err != nil {
    t.Error("Should be able to create file.")
  }
}

// Test to not create a file
func TestCreateFail(t *testing.T) {
  err := CreateFile("bouya/bouya.txt")

  if err == nil {
    t.Error("Should not be able to create file.")
  }
}

// Test to read the content of a file
func TestReadFile1(t *testing.T) {
  bytes, err := ReadFile("test/file1.txt")

  content := string(bytes)
  shouldBe := `This is a file.

This is a {pattern}.

This is not a {pattern.`

  if content != shouldBe || err != nil {
    t.Error("Could not read file.")
  }
}

// Test to try to read the content of an non-existing file
func TestReadFileFail(t *testing.T) {
  _, err := ReadFile("bouya")

  if err.Error() != "Error reading file bouya: file do not exist" {
    t.Error("Should not be able to read any content.")
  }
}

func TestWriteFile2(t *testing.T) {
  err := WriteFile("test/file2.txt", "Having fun")
}