package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf(
		"Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s",
		fname,
		lname,
		age,
		messageRate,
		isSubscribed,
		message,
	)

	// Don't touch below this line

	fmt.Println(userLog)
}

/* 
The function `fmt.Sprintf` formats a string according to the format specifier you've given it.
Then, it returns the formatted string without printing it to the console.
To display the string, you can use it with a print function, like `fmt.Println`,
to output the formatted string to the console.
It sort of prepares the string for you, and then you decide how and when to display it.

TIPS:
	fmt.Sprintf can be used to format strings.
	%.1f rounds a float to the tenths place, %.2f rounds to the hundredths place, etc.
	%t formats a boolean value.
	%v can be used to format any value in its default representation.
	%s can be used to format a string.
	%d can be used to format an integer.
 */
