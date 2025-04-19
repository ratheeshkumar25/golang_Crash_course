package main

import "fmt"

// 1. Basic Data Types:
// These are the fundamental types provided by Go and include:

// Numeric types:

// int, int8, int16, int32, int64

// uint, uint8, uint16, uint32, uint64

// float32, float64

// complex64, complex128

// String type: string

// Boolean type: bool

// Byte and rune types (aliases):

// byte (alias for uint8)

// rune (alias for int32)

func main() {

	// These are the basic data types in Golang for beginners to understand.
	// In Golang, there are two types of variable declarations: long and short variable declarations.

	// Long variable declaration: Declare and initialize a variable with an integer value.
	var a int
	a = 10
	fmt.Println("a value is", a)

	// Short variable declaration: Declare and initialize a variable with a float (decimal) value.
	b := 11.50
	fmt.Println("b value is", b)

	// String type example:
	// Long variable declaration
	var word string
	word = "Hi, I am Ratheesh"
	fmt.Println(word)

	// Short variable declaration
	word1 := "Welcome to the world of Go programming"
	fmt.Println(word1)

	// Boolean type example:
	// Long variable declaration: By default, a boolean variable is initialized to false.
	var ok bool
	fmt.Println("ok is", ok)

	// Short variable declaration: Assign a true value to the boolean variable.
	exist := true
	fmt.Println("exist is", exist)
}
