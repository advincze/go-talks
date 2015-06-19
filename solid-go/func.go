package main

import "log"

func main() {

}

type Object struct {
	Name  string
	Value int
}

// BEGIN OMIT

// parameter passed by value or reference
func secretFunc(a, b string, safeOutside Object, canBeChanged *Object) {

	// support for closures
	var addOne = func(x int) int {
		return x + 1
	}

	safeOutside.Value = addOne(canBeChanged.Value)

	canBeChanged.Name = "has been changed for the caller"
	safeOutside.Name = "change is only local"
}

//Access by first letter upper or lower case
func PublicAPIFunc() {
	log.Println("i am available for everyone")
}

// END OMIT
