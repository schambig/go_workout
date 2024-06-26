package main

import "fmt"

func main() {
	const name = "Salomon Chambi"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
}

/* 
Go follows the printf tradition from the C language.
It is not as elegant as Python's f-strings, unfortunately.

	fmt.Printf - Prints a formatted string to standard out.
	fmt.Sprintf() - Returns the formatted string

These following "formatting verbs" work with the formatting functions above:

DEFAULT REPRESENTATION
The %v variant prints the Go syntax representation of a value, it's a nice default.
	s := fmt.Sprintf("I am %v years old", 10)
	// I am 10 years old

	s := fmt.Sprintf("I am %v years old", "way too many")
	// I am way too many years old

If you want to print in a more specific way, you can use the following formatting verbs:

STRING
	s := fmt.Sprintf("I am %s years old", "way too many")
	// I am way too many years old
INTEGER
	s := fmt.Sprintf("I am %d years old", 10)
	// I am 10 years old
FLOAT
	s := fmt.Sprintf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	s := fmt.Sprintf("I am %.2f years old", 10.523)
	// I am 10.52 years old
 */
