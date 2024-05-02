package main

import "fmt"

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 10000
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	} else {
		return 0
	}
}

func main() {
    result := getMonthlyPrice("basic")
    result2 := getMonthlyPrice("premium")
    result3 := getMonthlyPrice("enterprise")
    
    fmt.Println(result, result2, result3)
}
