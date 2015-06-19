package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	obj := struct {
		Message string
	}{"Go!"}

	// BEGIN OMIT
	var i int
	i = 4

	var j = 8

	k := 15 // inferred type int
	// k = int64(99) //compile error : no implicit type conversion
	// k = "go"      //compile error

	l, m := 16, 23                  //multiple assignments
	bytes, err := json.Marshal(obj) //also used for error handling

	m, l = l, m //switch variables in one line

	// END OMIT

	fmt.Printf("i: %d\n", i)
	fmt.Printf("j: %d\n", j)
	fmt.Printf("k: %d\n", k)
	fmt.Printf("l: %d\n", l)
	fmt.Printf("m: %d\n", m)
	fmt.Printf("bytes: %s %v\n", bytes, err)

}
