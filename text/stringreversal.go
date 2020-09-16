package text

import (
	"fmt"
)

func reverseString(original string) string {
	var reversed string

	for index := (len(original) - 1); index >= 0; index-- {
		reversed += string(original[index])
	}
	return reversed
}

// ReverseString : Enter a string and the program will reverse it and print it out.
func ReverseString() {
	var userinput string
	fmt.Print("Enter string: ")
	fmt.Scan(&userinput)
	fmt.Println(reverseString(userinput))
}
