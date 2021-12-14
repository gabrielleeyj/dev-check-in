package main

import (
	"os"
	"testing"
)

var (
	root  = "./mock"
	files []os.DirEntry
	err   error
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
