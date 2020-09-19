package text

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// WordCounter :  Counts the number of individual words in a string.
// For added complexity read these strings in from a text file and generate a summary.
func WordCounter() {
	var inputType int
	fmt.Println("Press 1. To read from a file.")
	fmt.Println("Press 2. To enter a string.")
	fmt.Scan(&inputType)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string/file: ")
	if !scanner.Scan() {
		os.Exit(1)
	}
	text := scanner.Text()
	if inputType == 1 {
		file, _ := ioutil.ReadFile(text)
		text = string(file)
	}
	fmt.Println("Number of words(space-separated) is:", len(strings.Split(text, " ")))

}
