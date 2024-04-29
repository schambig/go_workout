package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}

/* 
Constants must be known at compile time. They are usually declared with a static value:
	const myInt = 15

However, constants can be computed as long as the computation can happen at compile time.
For example, this is valid:
	const firstName = "Salomon"
	const lastName = "Chambi"
	const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time
like you can in JavaScript. This breaks:
	// the current time can only be known when the program is running
	const currentTime = time.Now()
 */
