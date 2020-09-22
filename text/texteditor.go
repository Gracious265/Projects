package text

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GoTextEditor : Notepad style application that can open, edit, and save text documents.
// Optional: Add syntax highlighting and other features.
func GoTextEditor() {
	var filename, operation string

	fmt.Println("Enter operation (read/write): ")
	fmt.Scan(&operation)

	fmt.Println("Enter file name:")
	fmt.Scan(&filename)
	filename = strings.TrimSpace(filename)

	if operation == "read" {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(string(file))
	} else {
		var err error
		var file *os.File
		file, err = os.Create(filename)
		defer file.Close()

		if err != nil {
			os.Exit(1)
		}
		fmt.Println("Enter data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		data := scanner.Text()
		file.WriteString(data)
	}

}
