package numbers

import (
	"fmt"
	"math"
)

func contains(num int, list []int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

// GetPrimeFactor : Get prime factors of a number
func GetPrimeFactor() {
	var number int
	var factors []int
	fmt.Println("Enter a number to find it's prime factors: ")
	fmt.Scan(&number)
	for number%2 == 0 {
		if !contains(2, factors) {
			factors = append(factors, 2)
		}
		number = number / 2
	}
	for i := 3; i <= int(math.Sqrt(float64(number)))+1; {
		for number%i == 0 {
			number = number / i
			if !contains(i, factors) {
				factors = append(factors, i)
			}
		}
		i += 2
	}
	if !contains(number, factors) && number > 2 {
		factors = append(factors, 2)
	}
	fmt.Println(factors)
}
