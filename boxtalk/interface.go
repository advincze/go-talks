package main

import (
	"fmt"
)

// BEGIN INTERFACE OMIT
type Surface interface {
	Area() int
}

var s Surface = Rectangle{4, 4}

// structural typing: types implicitly fulfill interfaces
// BEGIN EMBEDDING OMIT
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

// END INTERFACE OMIT
type Square struct {
	Rectangle // Embedding
}

func NewSquare(length int) *Square {
	return &Square{Rectangle{length, length}}
}

func main() {
	fmt.Printf("square area: %d \n", NewSquare(4).Area())
}

// END EMBEDDING OMIT

//works not oly for structs
type OneDimensionalPoint int

func (o *OneDimensionalPoint) Area() int {
	return 0
}
