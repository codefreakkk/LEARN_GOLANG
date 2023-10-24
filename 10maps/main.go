package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["a"] = 1
	mp["a"]++
	mp["b"] = 2
	mp["c"] = 3

	// delete perticular key from map
	delete(mp, "b")

	fmt.Println(mp["a"])

	// declaration and initilization
	var m map[string]int = make(map[string]int)
	m["key"] = 1

	fmt.Println(m["key"])

	// basic for loop
	for key, value := range mp {
		fmt.Printf("Key -> %v, value -> %v \n", key, value)
	}
}
