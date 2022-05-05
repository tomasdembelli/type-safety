package main

import "fmt"

// addTwoThings expects 2 int type arguments, and add them.
func addTwoThings(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	addTwoThings(1, 2)

	// following calls will prevent program to be compiled.
	addTwoThings(1, 2.3)
	addTwoThings("a", "b")
	addTwoThings(1, "2")
}