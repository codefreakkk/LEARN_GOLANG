package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

/*
NOTE
 - Defer will execute at the last of function
 - if there are multiple defer stmt then it will execute in manner of LIFO Principle
*/
