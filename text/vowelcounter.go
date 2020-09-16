package text

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isVowel(character string) bool {
	var isvowel bool
	var vowels = [5]string{"a", "e", "i", "o", "u"}
	for _, value := range vowels {
		value = string(value)
		if character == value || character == strings.ToUpper(value) {
			isvowel = true
		}
	}
	return isvowel
}

// VowelCounter : Enter a string and the program counts the number of vowels in the text.
// For added complexity have it report a sum of each vowel found.
func VowelCounter() {
	var userinput string
	var vowelCounter, consonantCounter int

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string: ")
	userinput, _ = in.ReadString('\n')

	for _, char := range userinput {
		if isVowel(string(char)) {
			vowelCounter++
		} else {
			consonantCounter++
		}
	}
	fmt.Println("Vowels:", vowelCounter, "Consonants:", consonantCounter)

}
