package text

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// RegexQueryTool :  tool that allows the user to enter a text string and then in a separate control enter a regex pattern.
// It will run the regular expression against the source text and return any matches or flag errors in the regular expression.
func RegexQueryTool() {
	var text, pattern string
	fmt.Println("Enter pattern:")
	fmt.Scan(&pattern)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text string: ")
	if !scanner.Scan() {
		os.Exit(1)
	}
	text = scanner.Text()
	compiler, _ := regexp.Compile(pattern)
	matches := compiler.FindAllString(text, -1)
	length := len(matches)
	if length == 0 {
		fmt.Println("No matches found.")
	} else {
		fmt.Println(length, "matches found. ")
		for _, char := range matches {
			fmt.Print(char, " ")
		}
	}
}
