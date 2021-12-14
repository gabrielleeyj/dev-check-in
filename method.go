package main

import (
	"bufio"
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
)

func readDirectory(root string) ([]string, error) {
	var findList []string
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		findList = append(findList, path)
		return nil
	})
	return findList, err
}

func readChecker(path, keyword string) (string, error) {
	data, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	line := 1
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), []byte(keyword)) {
			return path, nil
		}
		line++
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return "", err
}
