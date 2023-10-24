package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	// for loop type 1
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// for loop type 2
	for i := range nums {
		fmt.Println(nums[i])
	}

	// for each loop
	for _, value := range nums {

		// case for go stmt
		if value == 2 {
			goto label
		}
		fmt.Println(value)
	}

	// go to label
label:
	fmt.Println("hiii")
}
