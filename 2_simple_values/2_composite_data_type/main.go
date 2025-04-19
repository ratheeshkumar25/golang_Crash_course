package main

import "fmt"

func main() {

	// Composite types
	fmt.Println("\nComposite types:")

	// Arrays in Go are fixed in size. An array stores elements of the same type
	// in contiguous memory locations.
	arr := [3]int{1, 2, 3}

	// A slice in Go is built on top of an array. It is essentially a dynamic,
	// resizable version of an array. If you modify the underlying array, the
	// changes will be reflected in the slice.
	slice := []int{4, 5, 6}

	// A map is an inbuilt data structure in Go. It stores data in key-value pairs.
	m := map[string]int{"one": 1, "two": 2}

	// A struct is a custom data type that can hold fields of different types.
	// It is declared using the 'type' keyword.
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "John", Age: 30} // Struct initialization

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", p)
}
