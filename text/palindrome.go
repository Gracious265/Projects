package text

import (
	"fmt"
)

// CheckPalindrome : Checks if the string entered by the user is a palindrome.
// That is that it reads the same forwards as backwards like “racecar”
func CheckPalindrome() {
	var userinput string
	fmt.Print("Enter string: ")
	fmt.Scan(&userinput)

	if userinput == reverseString(userinput) {
		fmt.Println("Palindrome.")
	} else {
		fmt.Println("Not palindrome.")
	}
}
