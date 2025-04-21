package main

import "fmt"

// Constants are values that cannot be changed after they are defined.
// They are useful for defining fixed values that remain constant throughout the program.
// Constants can be declared outside or inside functions.
// In this example given name of constant is not as per standar policy-Instead of generic names like a, b, and c, use descriptive names for constants

// Defining constants
// // An integer constant
const a = 10

// A string constant
const b = "Age"

// A boolean constant
const c = false

func main() {
	fmt.Println("age", a)
	fmt.Printf("%s  is %d", b, a)
	fmt.Printf("%d is less than %t", a, c)
}
