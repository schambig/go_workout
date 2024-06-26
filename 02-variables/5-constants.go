package main

import "fmt"

func main() {
	const premiumPlanName = "Premium Plan"
	// premiumPlanName = "Basic Plan" // can't change because of `const` so create a new var
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}

/* 
Constants are declared with the const keyword.They can't use the `:=` short declaration syntax.
	const pi = 3.14159

Constants can be character, string, boolean, or numeric values.
They can not be more complex types like slices, maps and structs, which are types we will study later.

As the name implies, the value of a constant can't be changed after it has been declared.
 */
