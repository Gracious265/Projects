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
		intNum := int(num - '0')
		sum += intNum * intNum
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
	var (
		isHappy bool
		limit   int
	)

	fmt.Print("Enter the limit: ")
	fmt.Scan(&limit)

	fmt.Print("Happy numbers are:\n1 10 ")
	count := 3
	for number := 11; count <= limit; number++ {
		isHappy = isHappyNumber(number)
		if isHappy {
			fmt.Printf("%d ", number)
			count++
		}
	}
}
