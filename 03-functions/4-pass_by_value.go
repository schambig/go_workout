package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	var bill int 
	bill = costPerSend * messagesSent
	return bill
}

/*
// alternative 
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}
 */

/* 
 Variables in Go are passed by value (except for a few data types we haven't covered yet).
 "Pass by value" means that when a variable is passed into a function,
 that function receives a copy of the variable.
 The function is unable to mutate the caller's original data.

	func main() {
		x := 5
		increment(x)
	
		fmt.Println(x)
		// still prints 5,
		// because the increment function
		// received a copy of x
	}
	
	func increment(x int) {
		x++
	}
 */
 