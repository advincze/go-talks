package main

func main() {

}

// BEGIN OMIT
type CustomerId int

type Article string

type Person struct {
	FirstName string
	LastName  string
	Age       int             //exported === all holders of Person can access
	secret    map[string]bool // private === only same package can access
}

// pointers (but no pointer arithmetic)
func (p *Person) RememberSecret(msg string) {
	p.secret[msg] = true
}

// END OMIT
