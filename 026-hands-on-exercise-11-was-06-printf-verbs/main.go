package main

import "fmt"

/*
Task:
for a VARIABLE storing a VALUE of TYPE
	string
	int
	float64
use print verbs to show
	value
	type
*/

func main() {
	s, i, f := "Happiness", 42, 42.42
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
	fmt.Printf("%v is of type %T\n", s, s)
}
