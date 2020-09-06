package number

import (
	"fmt"
	"strings"
)


// CreditCardValidator : Takes in a credit card number from a common credit card vendor (Visa, MasterCard, American Express, Discoverer) 
// and validates it to make sure that it is a valid number (look into how credit cards use a checksum).
func CreditCardValidator(){
	var cardNumber string
	var lastDigit int
	fmt.Print("Enter card number: ")
	fmt.Scan(&cardNumber)
	// TODO: https://www.freeformatter.com/credit-card-number-generator-validator.html#howToValidate
}