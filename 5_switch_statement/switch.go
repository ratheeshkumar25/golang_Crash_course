package main

import (
	"fmt"
	"time"
)

//A switch statement is a control structure in Go (and many other programming languages) used for conditional execution based on the value of a variable or expression.
//t allows you to match a value against multiple predefined cases and execute the corresponding block of code.
//

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	//use the map to updated values as key value pairs
	days := map[int]string{
		Sunday:    "Sunday",
		Monday:    "Monday",
		Tuesday:   "Tuesday",
		Wednesday: "Wednesday",
		Thursday:  "Thursday",
		Friday:    "Friday",
		Saturday:  "Saturday",
	}

	for i := 1; i <= 7; i++ {
		switch i {
		case Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday:
			// Fetch and print the string representation
			fmt.Printf("Day %d: %s\n", i, days[i])
		default:
			fmt.Println("Invalid day!")
		}
	}

	//Another example
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("Its working day")
	}

}
