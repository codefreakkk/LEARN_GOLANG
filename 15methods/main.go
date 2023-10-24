package main

import "fmt"

func main() {
	user := User{"Harsh said", "harshsaid31@gmail.com", 21}

	user.GetName()
	user.NewEmail()

	fmt.Println("Original email ", user.Email)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetName() {
	fmt.Println("Name of user is ", u.Name)
}

// by using pointers the original value will be changed
func (u *User) NewEmail() {
	u.Email = "harsh@gmail.com"
	fmt.Println("New Email is ", u.Email)
}
