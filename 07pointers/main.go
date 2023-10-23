package main

import "fmt"

func main() {
	num := 10
	var ptr = &num

	fmt.Println(*ptr)

	*ptr = *ptr + 2
	fmt.Println(*ptr)
}
