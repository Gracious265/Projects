package numbers

import (
	"math"
	"fmt"
)

func piELimit() {
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