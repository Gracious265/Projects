package number

import (
	"fmt"
	"strconv"
)

type units struct {
	unit  string
	teens string
	ty    string
}

var spells = map[int]units{
	1: units{unit: "one", teens: "eleven", ty: "ten"},
	2: units{unit: "two", teens: "twelve", ty: "twenty"},
	3: units{unit: "three", teens: "thirteen", ty: "thirty"},
	4: units{unit: "four", teens: "fourteen", ty: "forty"},
	5: units{unit: "five", teens: "fifteen", ty: "fifty"},
	6: units{unit: "six", teens: "sixteen", ty: "sixty"},
	7: units{unit: "seven", teens: "seventeen", ty: "seventy"},
	8: units{unit: "eight", teens: "eighteen", ty: "eighty"},
	9: units{unit: "nine", teens: "nineteen", ty: "ninety"},
}

const (
	teen = "teen"
	ty   = "ty"
	hund = "hundred"
	thou = "thousand"
)

// SpellNumbers : Show how to spell out a number in English.
// You can use a preexisting implementation or roll your own, but you should support inputs up to at least one million
// (or the maximum value of your language's default bounded integer type, if that's less).
// Optional: Support for inputs other than positive integers (like zero, negative integers, and floating-point numbers).
func SpellNumbers() string{
	var (
		number        int
		numeralString string
		ling          string
		length        int
		digit int
	)
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	numeralString = strconv.Itoa(number)
	length = len(numeralString)
	digit = int(numeralString[0] - '0')
	if length == 4 {
		ling = spells[digit].unit + " " + thou + " "
		numeralString = numeralString[1:]
		if numeralString == "000"{
			return ling
		}
		digit = int(numeralString[0] - '0')
		length--
	}

	if length == 3 {
		ling += spells[digit].unit + " " + hund + " "
		numeralString = numeralString[1:]
		digit = int(numeralString[0] - '0')
		length--
	}
	if length == 2 {
		if digit == 1 {
			digit = int(numeralString[1] - '0')
			ling += spells[digit].teens
			return ling
		} else {
			ling += spells[digit].ty
			digit = int(numeralString[1] - '0')
			if digit != 0 {
				ling += " " + spells[digit].unit
			}
			return ling
		}
		numeralString = numeralString[1:]
		digit = int(numeralString[0] - '0')
		length--
	} else if length == 1 {
		ling = spells[number].unit
		return ling
	}
	return ling
}
