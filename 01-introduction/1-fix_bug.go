package main

import "fmt"

func main() {
	messagesFromSalo := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}
	numMessages := float64(len(messagesFromSalo))
	costPerMessage := .02

	// don't touch above this line

	// totalCost := costPerMessage + numMessages
	totalCost := costPerMessage * numMessages

	// don't touch below this line

	fmt.Printf("Salo spent %.2f on text messages today\n", totalCost)
}
