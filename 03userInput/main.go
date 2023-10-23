package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name")

	// take input
	input := bufio.NewReader(os.Stdin)

	// Comma ok syntax
	name, _ := input.ReadString('\n') // here _ means nothing for that section
	fmt.Println("Your name is ", name)
}
