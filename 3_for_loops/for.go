package main

import "fmt"

func main() {
	//Classic for loop for used programming

	//print numbers from 1 to 5
	for i := 0; i <= 5; i++ {
		fmt.Println("Numbers", i)
	}

	//while loop example , in go for while implementation we are using for key word
	count := 1
	for count <= 3 {
		fmt.Println("Counters", count)
		// if count == 2 {
		// 	continue
		// }
		count++
	}

	//infinite loop in golang

	for {
		fmt.Println("Hi I am working us infinite loop")
		if count == 1 {
			break
		}
		break
	}

	//Range loop its also classic example for the for loops

	for counts := range count {

		if counts == 3 {
			fmt.Println("Count is equal to ", counts)
		}

	}

}
