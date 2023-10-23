package main

import "fmt"

// GLOBAL variable
var globalVar1 = "harsh"
var globalVar2 string = "harsh"

// CONST
// Note - since 'L' is captial LoginToken is considered as public variable
const LoginToken string = "sjdfhefsgfjsdfkjhsdgfkjh"

func main() {
	var username string = "harsh"
	username = "harsh sachin said"
	fmt.Printf("User name %s", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var smallValue uint8 = 33
	fmt.Println(smallValue)

	var smallFloatValue float32 = 45.22
	fmt.Println(smallFloatValue)

	// implicit type
	var name = "harsh"
	fmt.Println(name)

	// no var style
	numberOfUsers := 200
	fmt.Println(numberOfUsers)

	// Printing global variable
	fmt.Println(globalVar1)
	fmt.Println(globalVar2)
}
