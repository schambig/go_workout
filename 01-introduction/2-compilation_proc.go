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

/* 
- If you're developing and testing your Go script, using `go run <script>` is more convenient.
  It compiles and runs your script in one step, making it easier for quick iterations.

- If you've finished developing and are deploying your script or need it to run efficiently,
  compiling it first with `go build <script>` and then running the compiled binary with `./<script>` is recommended.
  This pre-compiled binary tends to start faster and doesn't require the Go toolchain to execute,
  making it more suitable for production environments.

Each method has its place. For development and testing, use `go run`.
For deployment, compile first and then run the binary.
 */
