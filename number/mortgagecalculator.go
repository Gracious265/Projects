package number

import (
	"fmt"
	"math"
)

// CalculateMortgage : Calculate the monthly payments of a fixed term mortgage  over given Nth terms at a given interest rate. Also figure out how long it will take the user to pay back the loan.
func CalculateMortgage() {
	var months int
	var loan, rate, payment, monthlyRate float64
	fmt.Print("Enter mortgage term (in months): ")
	fmt.Scan(&months)
	fmt.Print("Enter interest rate: ")
	fmt.Scan(&rate)
	fmt.Print("Enter loan value: ")
	fmt.Scan(&loan)
	monthlyRate = float64(rate / 100 / 12)
	payment = (monthlyRate / (1 - math.Pow((1+monthlyRate), -float64(months)))) * loan
	fmt.Printf("Monthly payment for a %.2f amount for %d months at a %.2f interest rate is: %f", loan, months, rate, payment)
}
