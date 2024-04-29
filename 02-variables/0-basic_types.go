package main

import "fmt"

func main() {
	// initialize variables here
	// The value of a declared variable with no assignment will be its zero value
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("int: %v, float64: %.2f, bool: %v, string: %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

/* 
Go's basic variable types are:

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
		 // represents a Unicode code point

	float32 float64

	complex64 complex128

We talked about `strings` and `ints` previously,
and those two types should be fairly self-explanatory.

A `bool` is a boolean variable, meaning it has a value of true or false.

The floating point types `float32` and `float64` are used for numbers that are not integers
-- that is, they have digits to the right of the decimal place, such as 3.14159.
	The `float32` type uses 32 bits of precision,
	while the `float64` type uses 64 bits to be able to more precisely store more digits.
 */
