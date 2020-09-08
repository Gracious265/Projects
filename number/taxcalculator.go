package number

import (
	"fmt"
)

// Constants can be character, string, boolean, or numeric values.
const(
	standardDeduction = 50000
	section80CTD = 150000
	section80DMaxDSelf = 25000
	section80DMaxDParent = 50000
	section80TTAMaxD = 10000
)


type investement struct{
	section80C, section80D, section80TTA float64
}

// TaxCalculator : Asks the user to enter a cost and either a country or state tax. It then returns the tax plus the total cost with tax.
func TaxCalculator(){
	var basicSalary, hra, specialAllowance, otherIncome, grossTotalSalaryI, investment80C, taxableIncome, totalTax float64
	var taxRegime string //only take deductions if value is "old"
	if 0 < taxableIncome < 250000{
		totalTax = 0
	} else if 250000 < taxableIncome < 500000{
		fmt.Println()
		// TODO: fix this
		// https://cleartax.in/paytax/TaxCalculator
	} else {
		fmt.Println()
	}
}