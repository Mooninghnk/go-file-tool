package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createFile(name string, inside string) {
	file, ok := os.Create(name)
	if ok != nil {
		fmt.Println("file cant be created")
		return
	}
	defer file.Close()
	file.WriteString(inside)
	fmt.Println("file created")
}

func readFile(name string) (error string) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("can't read this file %s", name)
		return
	}
	return string(file)
}
func removeFile(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Println("cant delete that file")
		return
	}
}

func main() {
	var choice int
	var fileName string
	var fileInside string
	var toRead string
	var toDelete string
	for {
		fmt.Print("Pick an Option: 0 EXIT ~ 1 CREATEFILE ~ 2 READFILE ~ 3 REMOVE FILE: ")
		fmt.Scanf("%d", &choice)
		fmt.Println(choice)
		if choice == 0 {
			fmt.Println("bye :)")
			return
		}
		if choice == 1 {
			fmt.Print("File Name: ")
			fmt.Scanf("%s", &fileName)
			fmt.Print("Print To File: ")
			fmt.Scanf("%s", &fileInside)
			createFile(fileName, fileInside)
		}
		if choice == 2 {
			fmt.Print("Enter a file name: ")
			fmt.Scanf("%s", &toRead)
			fmt.Println(readFile(toRead))
		}
		if choice == 3 {
			fmt.Print("Enter a file name to delete: ")
			fmt.Scanf("%s", &toDelete)
			removeFile(toDelete)
		}
	}
}
