package number

import (
	"fmt"
)

// ReturnChange : The user enters a cost and then the amount of money given. The program will figure out the change and the number of quarters, dimes, nickels, pennies needed for the change.
func ReturnChange() {
	var billingAmt, givenAmt, returnAmt int
	notesMap:= make(map[string]int)
	fmt.Print("Enter the billing amount: ")
	fmt.Scan(&billingAmt)
	fmt.Print("Enter the given amount: ")
	fmt.Scan(&givenAmt)

	returnAmt = givenAmt - billingAmt

	if returnAmt < 0 {
		fmt.Errorf("paid less than the billing amount")
	} else{
		fmt.Printf("Return amount: %d \n", returnAmt)
	}

	for returnAmt > 2 {
		if returnAmt >= 2000 {
			notesMap["2000"] = (returnAmt - (returnAmt % 2000)) / 2000
			returnAmt %= 2000
		}
		if returnAmt >= 500 {
			notesMap["500"] = (returnAmt - (returnAmt % 500)) / 500
			returnAmt %= 500
		}
		if returnAmt >= 200 {
			notesMap["200"] = (returnAmt - (returnAmt % 200)) / 200
			returnAmt %= 200
		}
		if returnAmt >= 100 {
			notesMap["100"] = (returnAmt - (returnAmt % 100)) / 100
			returnAmt %= 100
		}
		if returnAmt >= 50 {
			notesMap["50"] = (returnAmt - (returnAmt % 50)) / 50
			returnAmt %= 50
		}
		if returnAmt >= 20 {
			notesMap["20"] = (returnAmt - (returnAmt % 20)) / 20
			returnAmt %= 20
		}
		if returnAmt >= 10 {
			notesMap["10"] = (returnAmt - (returnAmt % 10)) / 10
			returnAmt %= 10
		}
		if returnAmt >= 5 {
			notesMap["5"] = (returnAmt - (returnAmt % 5)) / 5
			returnAmt %= 5
		}
		if returnAmt >= 2 {
			notesMap["2"] = (returnAmt - (returnAmt % 2)) / 2
			returnAmt %= 2
		}
	}
	notesMap["1"] = returnAmt
	for key, value := range notesMap {
		if value > 0{
		fmt.Printf("%s rupee note/coin: %d \n", key, value)
		}
	}
}
