package main

import "fmt"

func main() {
	//simple way to understand if , else if and else statement
	ageOfStudent := 16

	if ageOfStudent >= 18 {
		fmt.Println("Quailified for selection")
	} else if ageOfStudent < 18 && ageOfStudent >= 15 {
		fmt.Println("Attend exam and prepare for the prequalification")
	} else {
		fmt.Println("Student is not qualified wait for the age criterial to meet")
	}

	role := "user"
	hasPermission := true

	if role == "user" && hasPermission {
		fmt.Println("Welcome to student world")
	}

	if age := 18; age >= 18 {
		fmt.Println("Full Access granted")
	} else {
		fmt.Println("Limited Access")
	}
}
