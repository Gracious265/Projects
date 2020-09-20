package text

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GoTextEditor : Notepad style application that can open, edit, and save text documents.
// Optional: Add syntax highlighting and other features.
func GoTextEditor() {
	var filename, operation string
	fmt.Println("Enter operation (read/edit/write): ")
	fmt.Scan(&operation)
	fmt.Println("Enter file name:")
	fmt.Scan(&filename)

	file, error := ioutil.ReadFile(filename)
	if error != nil {
		panic(error)
	}
	if operation == "read" {
		fmt.Println(string(file))
	}
	if operation == "write" {
		var data []byte
		fmt.Println("Enter data: ")
		fmt.Scan(&data)
		err := ioutil.WriteFile(filename, data, 0777)
		if err != nil {
			panic(error)
		}
		fmt.Println("Data printed successfully.")
	}

	if operation == "edit" {
		var data []byte
		fmt.Println("Enter data: ")
		fmt.Scan(&data)
		file, error := os.OpenFile(filename, os.O_APPEND, 0777)
		defer file.Close()
		if error != nil {
			panic(error)
		}
		_, err := file.Write(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(ioutil.ReadFile(filename))

	}

}

// TODO: https://gobyexample.com/writing-files
