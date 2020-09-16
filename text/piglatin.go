package text

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PigLatin : Pig Latin is a game of alterations played on the English language game.
// To create the Pig Latin form of an English word the initial consonant sound is transposed
// to the end of the word and an ay is affixed (Ex.: "banana" would yield anana-bay).
func PigLatin() {
	var userinput string
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string: ")
	userinput, _ = in.ReadString('\n')
	for _, str := range strings.Split(userinput, " ") {
		firstCharacter := string(str[0])
		if !isVowel(firstCharacter) {
			str = strings.Replace(str, firstCharacter, "", 1)
			str = strings.Replace(str, "\n", "", -1)
			str += "-" + firstCharacter + "ay"
			fmt.Println(str)
		}
	}
}
