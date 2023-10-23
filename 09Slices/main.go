package main

import (
	"fmt"
	"sort"
)

func main() {

	// Slice syntax
	var nums = []int{}

	// append metod
	nums = append(nums, 1, 2, 3, 4, 5)

	fmt.Println(nums[1:])  // start from 0 skip 1 element and take rest of elements
	fmt.Println(nums[1:3]) // index 3 non-inclusive

	// make() method
	scores := make([]int, 2)
	scores[0] = 100
	scores[1] = 90

	// append into scores
	scores = append(scores, 89, 91)
	fmt.Println(scores)

	// sort slice
	sort.Ints(scores)
	fmt.Println(scores)

	// remove element by index from slice
	var numbers = []int{1, 2, 3, 4, 5, 6}
	var index int = 2
	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Println(numbers)

}
