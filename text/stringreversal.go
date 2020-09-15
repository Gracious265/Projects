package text

import (
	"fmt"
)

// ReverseString : Enter a string and the program will reverse it and print it out.
func ReverseString(){
	var userinput, reversed string
	fmt.Print("Enter string: ")
	fmt.Scan(&userinput)
	for index:=(len(userinput) - 1); index >= 0; index-- {
		fmt.Println(index)
		reversed += string(userinput[index])
	}
	fmt.Println(reversed)	
}