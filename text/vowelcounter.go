package text

import (
	"fmt"
)

func isVowel(character string) bool{
	var isvowel bool
	var vowels = [5]string{"a","e","i","o","u"}
	for _, value := range(vowels){
		if character == string(value){
			isvowel = true
		}
	}
	return isvowel
}

// VowelCounter : Enter a string and the program counts the number of vowels in the text.
// For added complexity have it report a sum of each vowel found.
func VowelCounter(){
	var userinput string
	var vowelCounter, consonantCounter int

	fmt.Print("Enter string: ")
	// TODO: Figure out how to take multiple strings as input.
	fmt.Scanf("%s", &userinput)

	for _, char := range userinput {
		if isVowel(string(char)){
			vowelCounter++
		} else {
			consonantCounter++
		}
	}
	fmt.Println("Vowels:",vowelCounter, "Consonants:",consonantCounter)
	
}