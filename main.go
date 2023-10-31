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

	// Formating output
	nameoutput := fmt.Sprintf("Your full name is %s %s", name, lastname)

	// Printing full name
	fmt.Println("")
	fmt.Println(nameoutput)
	fmt.Println("")

	// Invoking function "contact"
	contact()
}

func contact() {
	// Creating variable "email" as string and obtaining value of user input
	var email string
	fmt.Print("Enter your e-mail: ")
	fmt.Scanln(&email)

	// Creating variable "cellphone" as int and obtaining value of user input
	var cellphone int
	fmt.Print("Enter your cellphone number, including your country code: +")
	fmt.Scanln(&cellphone)

	// Formating output
	contactoutput := fmt.Sprintf("Your e-mail is %s and your cellphone number is +%d", email, cellphone)

	// Printing contact methods
	fmt.Println()
	fmt.Println(contactoutput)
	fmt.Println()
}
