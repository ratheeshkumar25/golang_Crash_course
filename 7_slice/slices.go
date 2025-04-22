package main

import (
	"fmt"
	"sort"
)

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

	//1.find the min val in the arry

	arr := []int{10, 20, 30, 50, 5, 55}

	minVal := func(arr []int) int {
		minV := arr[0]

		for _, v := range arr {
			if v < minV {
				minV = v
			}
		}
		return minV
	}
	fmt.Println("Minum value of array", minVal(arr))

	arr1 := []int{1, 1, 3, 3, 3, 10, 1, 0, 10, 0, 2, 2, 2}
	freq := freqOfArray(arr1)
	for _, pair := range freq {
		fmt.Printf("Value: %d, Frequency: %d\n", pair[0], pair[1])
	}
	//with optimal solution
	freqs := freqOfArrays(arr1)

	for i, v := range freqs {
		fmt.Printf("With map approach Value %d , frequecy %d \n", i, v)
	}

	//2.missing number
	arr2 := []int{0, 1, 5, 6, 4, 2}
	missingNum := missingNumber(arr2)
	fmt.Println("missing number is", missingNum)
}

// 1.find the frequency of array brute force approach
func freqOfArray(arr []int) [][]int {
	result := [][]int{}

	for i := 0; i < len(arr); i++ {
		count := 0

		alreadyVist := false

		for _, pair := range result {
			if pair[0] == arr[i] {
				alreadyVist = true
				break
			}

			if alreadyVist {
				continue
			}
		}

		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		result = append(result, []int{arr[i], count})
	}
	return result
}

// with map optimal solution
func freqOfArrays(arr []int) map[int]int {
	freqMap := make(map[int]int)

	for _, v := range arr {
		freqMap[v]++
	}
	return freqMap
}

//2.Missing Number

func missingNumber(arr []int) int {
	//sort the array u can use bubble sort or merge sort , here i have used inbuilt sorting function from go standard liabiray
	sort.Ints(arr)
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] != i {
			return i
		}
	}
	return n
}
