package number

import (
	"fmt"
)

// Calculate :  A simple calculator to do basic operators.
func Calculate() {
	var operation string
	fmt.Print("Enter an operation symbol(+, -, *, /): ")
	fmt.Scan(&operation)
	var num1, num2 float64
	fmt.Print("Enter number1: ")
	fmt.Scan(&num1)
	fmt.Print("Enter number2: ")
	fmt.Scan(&num2)
	switch operation {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)

	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println(num1 + num2)
	}
}
