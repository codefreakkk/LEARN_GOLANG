package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is test content for file."

	// create file
	file, err := os.Create("./temp.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Length of file is ", length)

	readFile("./temp.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))
}
