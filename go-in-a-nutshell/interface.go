package main

import (
	"fmt"
)

// BEGIN OMIT

type Shaper interface {
	Area() int
}

// Duck typing : types implicitly fulfill interfaces
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

// Embedding
type Square struct {
	Rectangle
}

func NewSquare(length int) *Square {
	return &Square{Rectangle: Rectangle{length, length}}
}

func main() {
	fmt.Printf("square area: %d \n", NewSquare(4).Area())
}

// END OMIT

//works not oly for structs
type OneDimensionalPoint int

func (o *OneDimensionalPoint) Area() int {
	return 0
}
