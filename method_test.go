package main

import (
	"bufio"
	"bytes"
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

// Test to check if I'm able to read the file from path
func TestReadFile(t *testing.T) {
	for _, path := range findList {
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(data)
	}
}

// Scan the file line by line to find the matching string "TODO"
func TestReadScan(t *testing.T) {
	var todoPaths []string
	for _, path := range findList {
		data, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(data)
		line := 1
		for scanner.Scan() {
			if bytes.Contains(scanner.Bytes(), []byte("TODO")) {
				todoPaths = append(todoPaths, path)
			}
			line++
		}

		err = scanner.Err()
		if err != nil {
			panic(err)
		}
	}

	for _, path := range todoPaths {
		println(path)
	}
}
