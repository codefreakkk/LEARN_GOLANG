package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter a number")

	// create reader
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Print(num + 1)
	} else {
		fmt.Println(err)
	}
}
