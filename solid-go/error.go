package main

import "os"

func main() {

	filename := "myfile.txt"
	// BEGIN ERROR OMIT
	file, err := os.Create(filename)
	if err == nil {
		return
	}
	// END ERROR OMIT
	println(file)
}

// BEGIN EXCEPTION OMIT

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}

// END EXCEPTION OMIT
