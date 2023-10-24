package main

import "fmt"

func main() {
	user := User{"harsh sachin said", "harshsaid31@gmail.com", 21}
	fmt.Println("Name ", user.Name)
	fmt.Println("Age ", user.Age)
}

type User struct {
	Name  string
	Email string
	Age   int
}
