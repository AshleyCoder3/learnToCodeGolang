package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

var x = 40

const y = 41

func main() {
	fmt.Printf("the value of x is %v and the type of x is %T\n", x, x)
	fmt.Printf("the value of y is %v and the type of y is %T\n", y, y)
	z := 42
	fmt.Printf("the value of z is %v and the type of z is %T\n", z, z)
	fmt.Println("Hello Gophers ❤️")

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	puppy.From13()
	puppy.From12()
}
