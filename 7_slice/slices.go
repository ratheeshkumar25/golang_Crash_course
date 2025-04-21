package main

import "fmt"

// A slice in Go is built on top of an array. It is essentially a dynamic,
// resizable version of an array. If you modify the underlying array, the
// changes will be reflected in the slice.

func main() {
	// Empty slice
	slic := []int{} // An empty slice (length: 0, capacity: 0)
	// Nil slice
	var slcs []int // A nil slice (no memory allocated)

	fmt.Println("Empty slice:", slic, "Length:", len(slic), "Capacity:", cap(slic))
	fmt.Println("Nil slice:", slcs, "Length:", len(slcs), "Capacity:", cap(slcs))

	// Memory allocation using make
	nums := make([]int, 4, 5) // Slice with length 4 and capacity 5
	fmt.Println("Slice with make:", nums, "Length:", len(nums), "Capacity:", cap(nums))

	// Modifying a slice
	nums[0] = 10
	nums[1] = 20
	fmt.Println("Modified slice:", nums)

	// Appending an element (still within capacity)
	nums = append(nums, 30)
	fmt.Println("After appending:", nums, "Length:", len(nums), "Capacity:", cap(nums))

	// Appending beyond capacity (forces reallocation)
	nums = append(nums, 40)
	fmt.Println("After exceeding capacity:", nums, "Length:", len(nums), "Capacity:", cap(nums))

	//copy the slice
	num := make([]int, len(nums))
	copy(num, nums)
	fmt.Println("after copying slice", num)
}
