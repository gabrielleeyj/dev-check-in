package main

// setting: test directory location
var root = "./mock"

func main() {

	// run method to get list of directory
	dir, err := readDirectory(root)
	if err != nil {
		println(err)
	}

	// traverse through the list of directory and read the file
	for _, path := range dir {
		data, err := readChecker(path)
		if err != nil {
			println(err)
		}

		// if readChecker returns valid append to verified slice
		println(data)
	}
}
