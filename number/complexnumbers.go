package number

import (
	"fmt"
)

func add(num1, num2 complex64) complex64 {
	return num1 + num2
}
func multiplication(num1, num2 complex64) complex64 {
	return num1 * num2
}
func negation(num complex64) complex64 {
	return -(num)
}
func inversion(num complex64) complex64 {
	return 1 / num
}

func division(num1, num2 complex64) complex64 {
	return num1 * inversion(num2)
}
func substraction(num1, num2 complex64) complex64 {
	return num1 - num2
}

// ComplexNumbersCalculator : Show addition, multiplication, negation, and inversion of complex numbers in separate functions.
// (Subtraction and division operations can be made with pairs of these operations.) Print the results for each operation tested.
func ComplexNumbersCalculator() {
	var operation string
	var realNum, imagNum float32
	var num1, num2, result complex64
	fmt.Print("What operation would you like to perform?\nChoose from (addition, multiplication, negation, inversion, substraction, division): ")
	fmt.Scan(&operation)

	fmt.Print("Enter real and imaginary number: ")
	fmt.Scan(&realNum, &imagNum)
	num1 = complex(realNum, imagNum)

	if operation != "negation" && operation != "inversion" {
		fmt.Print("Enter another real and imaginary number: ")
		fmt.Scan(&realNum, &imagNum)
		num2 = complex(realNum, imagNum)
	}

	switch operation {
	case "substraction":
		result = substraction(num1, num2)
	case "multiplication":
		result = multiplication(num1, num2)
	case "inversion":
		result = inversion(num1)
	case "negation":
		result = negation(num1)
	case "division":
		result = division(num1, num2)
	default:
		result = add(num1, num2)
	}

	fmt.Println("Result:", result)
}
