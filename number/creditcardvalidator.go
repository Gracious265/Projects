package number

import (
	"fmt"
	"strconv"
)

// CreditCardValidator : Takes in a credit card number from a common credit card vendor (Visa, MasterCard, American Express, Discoverer)
// and validates it to make sure that it is a valid number (look into how credit cards use a checksum).
// Reference - https://www.freeformatter.com/credit-card-number-generator-validator.html#howToValidate
func CreditCardValidator() {
	var cardNumber, inverseCardNumber string
	var lastDigit, sum int
	fmt.Print("Enter card number: ")
	fmt.Scan(&cardNumber)
	for i := len(cardNumber) - 1; i > -1; i-- {
		if i == len(cardNumber)-1 {
			lastDigit, _ = strconv.Atoi(string(cardNumber[i]))
		} else {
			inverseCardNumber += string(cardNumber[i])
		}
	}
	for i := 0; i < len(inverseCardNumber); i++ {
		num, _ := strconv.Atoi(string(inverseCardNumber[i]))
		if (i+1)%2 != 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	fmt.Println("Sum ", sum)
	if sum%10 == lastDigit {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
