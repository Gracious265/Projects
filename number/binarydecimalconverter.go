package number

import (
	"fmt"
	"math"
)

func binaryToDecimal(number int) {
	var sum, prod float64
	var numberString string
	numberString = fmt.Sprint(number)
	for index, value := range numberString {
		if string(value) == "1" {
			prod = math.Pow(2, float64(index))
		} else {
			prod = 0
		}
		sum += prod
	}
	fmt.Println(sum)
}

func decimalToBinary(number int) {
	// Confirm the answers here: https://www.rapidtables.com/convert/number/decimal-to-binary.html
	var quotient, remainder int
	var binaries []int
	quotient = number / 2
	for quotient != 0 {
		quotient, remainder = number/2, number%2
		binaries = append(binaries, remainder)
		number = quotient
	}

	for i := len(binaries) - 1; i >= 0; i-- {
		fmt.Print(binaries[i])
	}
	fmt.Println()
}

// BinaryDecimalConversion : Develop a converter to convert a decimal number to binary or a binary number to its decimal equivalent.
func BinaryDecimalConversion() {
	var number int
	var conversionType string
	fmt.Print("Press B to convert from binary to decimal.\nPress D to convert from decimal to binary.\nEnter B/D: ")
	fmt.Scan(&conversionType)
	fmt.Print("\nEnter the number: ")
	fmt.Scan(&number)
	if conversionType == "D" {
		decimalToBinary(number)
	} else if conversionType == "B" {
		binaryToDecimal(number)
	} else {
		fmt.Errorf("wrong conversion value passed")
	}

}
