package main

import "fmt"

func main() {
	fmt.Println("the compiled textio server is starting")
}

/* 
`package main`
lets the Go compiler know that we want this code to compile and run as a standalone program,
as opposed to being a library that's imported by other programs.

`import fmt`
imports the fmt (formatting) package. The formatting package exists in Go's standard library
and lets us do things like print text to the console.

`func main()`
defines the main function. main is the name of the function that acts as the entry point for a Go program.
 */
