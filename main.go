package main

// setting: test directory location
var directory = "./mock"

func main() {

	// run method to get list of directory
	dir, err := readDirectory(directory)
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
		if data != "" {
			println("found keyword at: %v", data)
		}
	}
}
