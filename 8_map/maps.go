package main

import "fmt"

func main() {
	//long variable
	var Seen map[int]int

	fmt.Println("map", Seen)
	//Shor variable declaration
	//freq := make(map[int]int)

	//count the frequency of string

	s := "Hi Welcome to Golang Crash Course"

	strMap := make(map[rune]int)

	for _, v := range s {
		//avoid the space
		if v == ' ' {
			continue
		}
		//convert into small letter
		if v > 'A' && v < 'Z' {
			v = v + 32
		}
		strMap[v]++
	}

	fmt.Println("Frequency of string is")
	for word, count := range strMap {
		fmt.Printf("%c word %d count\n", word, count)
	}

}
