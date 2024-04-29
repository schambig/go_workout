package main

import "fmt"

func main() {
	// declare here, use walrus operator
	messageStart := "Happy birthday! You are now"
	age := 69
	messageEnd := "years old!"
	
	// don't edit below this line
	fmt.Println(messageStart, age, messageEnd)
}

/* 
Inside a function (like the main function) the := short assignment statement can be used
in place of a var declaration. The := operator infers the type of the new variable
based on the value.
It's colloquially called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:
	var empty string
	empty := ""

numCars := 10 // inferred as an integer
temperature := 0.0 // temperature is inferred as a float because it has a decimal
var isFunny = true // inferred as a boolean

Outside of a function (in the global/package scope), every statement begins with
a keyword (var, func, and so on) and so the := construct is not available.
 */
