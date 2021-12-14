# dev-check-in

Script in Golang using Go 1.6 std lib methods to traverse directories and read the files

Using only the standard library.

It will find all possible paths from root directory, then open all the files inside the directory and checks for the string "TODO"

## Instructions

After cloning the repository, to run the program:

1. At the root path of the cloned repository, run `go mod init dev-check-in` 
2. Then, `go run dev-check-in`
3. You should be able to see the returned results of the paths that have "TODO" inside.
