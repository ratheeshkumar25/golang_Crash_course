package main

import "fmt"

// Arrays in Go are fixed in size. An array stores elements of the same type
// in contiguous memory locations.

func main() {
	// array declared with length of 4
	var arr [4]int
	fmt.Println("before intialize element array", arr)
	//intialize value to array
	// arr[0] = 10
	// arr[1] = 11
	// arr[2] = 12
	// arr[3] = 14

	//find the array length we are using the len keyword in golang
	fmt.Println("length of array", len(arr))
	n := len(arr)
	//now adding value using loop"
	for i := 0; i < n; i++ {
		arr[i] = i + 1

	}
	fmt.Println("array of element is", arr)

	//twoD array
	//guess the output before run the program
	nums := [3][3]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("TwoD array", nums)

	// This code will result in a compile-time error because the array is defined with a size of 2,
	// but it is being initialized with 3 elements.
	// To fix this, either change the array size to match the number of elements (e.g., [3]int),
	// or use a slice (e.g., []int) which allows dynamic sizing.
	// num := [2]int{1, 2, 3} // This line will cause an error.

}
