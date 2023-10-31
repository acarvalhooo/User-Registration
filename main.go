package main

import "fmt"

func main() {
	// Creating variable "name" as string and obtaining value of user input
	var name string
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&name)

	// Creating variable "lastname" as string and obtaining value of user input
	var lastname string
	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastname)

	// Printing full name of user
	fmt.Println("Your full name is " + name + " " + lastname)
}
