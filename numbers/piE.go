package numbers

import (
	"fmt"
	"math"
)

// GetPiELimit : Enter a number and have the program generate PI or E up to that many decimal places. Keep a limit to how far the program will go.
func GetPiELimit() {
	var function, limit, printer string
	fmt.Println("Enter a number: ")
	fmt.Scan(&limit)
	fmt.Println("Enter the function (pi/E): ")
	fmt.Scan(&function)
	formatter := "%." + limit + "f"
	if function == "pi" {
		printer = fmt.Sprintf(formatter, math.Pi)
	} else if function == "E" {
		printer = fmt.Sprintf(formatter, math.E)
	} else {
		printer = "You did not enter a valid function."
	}
	fmt.Println(printer)
}
