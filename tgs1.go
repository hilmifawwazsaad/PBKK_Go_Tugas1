package main

import "fmt"

func main() {
	var input int
	fmt.Println("Welcome to Hallo Universe")
	fmt.Println("Please enter a number: ")
	fmt.Scan(&input)

	if input == 42 {
		fmt.Println("Hallo Universe!")
	} else {
		fmt.Println(input)
	}
}
