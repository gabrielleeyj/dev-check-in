package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var (
	root     = "./mock"
	files    []os.DirEntry
	err      error
	findList []string
)

// Checks the directories available from root
// only returns first level from root.
func TestReadDirectory(t *testing.T) {
	files, err = os.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		println(file.Name())
	}
}

// Able to check all possible paths of a directory
// Originally use Walk but the latest 1.6 documentation recommended WalkDir which is more efficient
func TestNestedDirectoryRead(t *testing.T) {
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// checks if it is a directory
		if info.IsDir() {
			return nil
		}

		// appends full path of files to findList
		findList = append(findList, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	// check to see if the paths are appended
	for _, path := range findList {
		println(path)
	}
}
