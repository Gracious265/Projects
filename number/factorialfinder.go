package number

import (
	"fmt"
)

func getFactorialRecursive(number int) int {
	if number > 0 {
		return number * getFactorialRecursive(number-1)
	}
	return 1
}

func getFactorialLooping(number int) int {
	var factorial int = 1
	for ; number > 1; number-- {
		factorial *= number
	}
	return factorial
}

// FactorialFinder : Finds factorial of a number.
func FactorialFinder() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Println("Factorial via recusion is:", getFactorialRecursive(num))
	fmt.Println("Factorial via looping is:", getFactorialLooping(num))
}
