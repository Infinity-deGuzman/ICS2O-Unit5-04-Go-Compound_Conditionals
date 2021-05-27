// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program shows what price you have to pay to get in the museum.

package main

import "fmt"

func main() {
	var userAge int
	var userDay string

	fmt.Println("Enter your age and the day of the week:")
	fmt.Println()
	fmt.Print("Age: ")
	fmt.Scanln(&userAge)
	fmt.Print("Day of the Week: ")
	fmt.Scanln(&userDay)

	if userAge < 5 || userAge > 95 {
		print("You have free admission :)")
	} else if (userAge > 12 && userAge < 21) || (userDay == "Tuesday" || userDay == "Thursday") {
		print("You are eligible for student pricing!")
	} else {
		print("You must pay regular pricing.")
	}
}
