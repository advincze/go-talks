package main

import (
	"fmt"
	"net/http"
)

// BEGIN OMIT
func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

// END OMIT
