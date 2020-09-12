package number

import (
	"fmt"
	"strconv"
)

func isHappyNumber(number int) bool {
	// TODO: Fix this
	var (
		sum    int
		strnum string
		happy  bool
	)
	strnum = strconv.Itoa(number)
	for _, num := range strnum {
		sum += int(num) * int(num)
	}
	if len(strconv.Itoa(sum)) > 1 {
		return isHappyNumber(sum)
	}
	if sum == 1 {
		happy = true
	} else {
		happy = false
	}

	return happy
}

// HappyNumbers : A happy number is defined by the following process.
// Starting with any positive integer, replace the number by the sum of the squares of its digits,
// and repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy numbers, while those that do not end in 1 are unhappy numbers.
// Display an example of your output here. Find first 8 happy numbers.
func HappyNumbers() {
	var happyNumbers = [8]int{1, 10}
	var (
		isHappy bool
	)
	index := 2
	for number := 11; index <= 7; number++ {
		isHappy = isHappyNumber(number)
		if isHappy {
			happyNumbers[index] = number
		}
		index++
	}
	fmt.Println("Happy numbers are: ", happyNumbers)
}
