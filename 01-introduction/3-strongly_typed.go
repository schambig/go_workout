package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	// var password int = 12345
	// change data type to string to be able to concatenate
	var password string = "12345"

	// don't edit below this line
	// line below will add spaces between each value when fmt.Println prints them
	// fmt.Println("Authorization: Basic", username, ":", password)
	fmt.Println("Authorization: Basic", username+":"+password)
}
