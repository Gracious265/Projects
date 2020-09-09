package number

import (
	"fmt"
)

// Constants can be character, string, boolean, or numeric values.
const (
	standardDeduction            = 50000
	section80CMaxDeduction       = 150000
	section80DMaxDeductionSelf   = 25000
	section80DMaxDeductionParent = 50000
	section80TTAMaxDeduction     = 10000
)

type investement struct {
	section80C, section80D, section80TTA float64
}

func getTotalTax(taxableIncome float64) float64 {
	var totalTax float64
	if taxableIncome >= 1500000 {
		totalTax = (taxableIncome-1500000)*0.3 + 187500
	}
	if taxableIncome >= 1250000 && taxableIncome < 1500000 {
		totalTax = (taxableIncome-1250000)*0.25 + 125000
	}
	if taxableIncome >= 1000000 && taxableIncome < 1250000 {
		totalTax = (taxableIncome-1000000)*0.2 + 75000
	}
	if taxableIncome >= 750000 && taxableIncome < 1000000 {
		totalTax = (taxableIncome-750000)*0.15 + 37500
	}
	if taxableIncome >= 500000 && taxableIncome < 750000 {
		totalTax = (taxableIncome-500000)*0.1 + 12500
	}
	if taxableIncome >= 250000 && taxableIncome < 500000 {
		totalTax = (taxableIncome - 250000) * 0.05
	}
	totalTax += totalTax * 0.04
	return totalTax
}

func calculateGrossTotalIncome(taxRegime string) float64 {
	var basicSalary, hra, hraExemption, lta, ltaExemption, specialAllowance, grossTotalSalaryI float64

	fmt.Print("Basic Salary: ")
	fmt.Scan(&basicSalary)
	fmt.Print("\nHRA: ")
	fmt.Scan(&hra)
	fmt.Print("\nLTA: ")
	fmt.Scanf("%f %f", &lta)
	fmt.Print("\nSpecial Allowance: ")
	fmt.Scan(&specialAllowance)

	grossTotalSalaryI = basicSalary + specialAllowance + hra

	if taxRegime == "new" {
		grossTotalSalaryI += lta
	} else if taxRegime == "old" {
		fmt.Print("\nHRA Exemption/Deduction: ")
		fmt.Scan(&hraExemption)
		fmt.Print("\nLTA Exemption/Deduction: ")
		fmt.Scan(&ltaExemption)
		grossTotalSalaryI -= standardDeduction
		grossTotalSalaryI -= hraExemption
		grossTotalSalaryI += lta - ltaExemption
	}
	return grossTotalSalaryI
}

func calculateTotalInvestment() float64 {
	var totalInvestment, investment80C, investment80TTA, investmentParent80D, investmentSelf80D float64
	fmt.Print("\nTotal investment under Section 80C: ")
	fmt.Scan(&investment80C)
	fmt.Print("\nTotal investment under Section 80D for Parent and Self: ")
	fmt.Scan(&investmentParent80D, &investmentSelf80D)
	fmt.Print("\nTotal investment under Section 80TTA: ")
	fmt.Scan(&investment80TTA)

	if investment80C > section80CMaxDeduction {
		totalInvestment += section80CMaxDeduction
	} else {
		totalInvestment += investment80C
	}
	if investmentParent80D > section80DMaxDeductionParent {
		totalInvestment += section80DMaxDeductionParent
	} else {
		totalInvestment += investmentParent80D
	}
	if investmentSelf80D > section80DMaxDeductionSelf {
		totalInvestment += section80DMaxDeductionSelf
	} else {
		totalInvestment += investmentSelf80D
	}
	if investment80TTA > section80TTAMaxDeduction {
		totalInvestment += section80TTAMaxDeduction
	} else {
		totalInvestment += investment80TTA
	}
	return totalInvestment
}

func getOtherInvestments() float64 {
	var otherIncome float64
	fmt.Print("\nEnter any other incomes in total: ")
	fmt.Scan(&otherIncome)
	return otherIncome
}

// TaxCalculator : Asks the user to enter a cost and either a country or state tax. It then returns the tax plus the total cost with tax.
// https://cleartax.in/paytax/TaxCalculator
func TaxCalculator() {
	var taxableIncome float64
	var taxRegime string
	fmt.Print("Enter tax regime, select from ('new', 'old'): ")
	fmt.Scan(&taxRegime)
	taxableIncome = calculateGrossTotalIncome(taxRegime) + getOtherInvestments()
	if taxRegime == "old" {
		taxableIncome -= calculateTotalInvestment()
	}
	fmt.Println("Your total tax is (it may not be accurate with old tax regime): ", getTotalTax(taxableIncome))
}
