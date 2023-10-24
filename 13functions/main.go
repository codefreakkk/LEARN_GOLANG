package main

import "fmt"

func main() {
	greet()

	result := add(1, 2)
	fmt.Println("Result ", result)

	sum := proAdder(1, 2, 3, 4)
	fmt.Println("Sum of all elements ", sum)

	add, msg := _2ReturnValue(1, 2)
	fmt.Printf("Addition is %d, msg is %s \n", add, msg)

	// getting array from function
	nums := getArray()
	fmt.Println(nums[1])

	// Anonymous function
	func() {
		fmt.Println("Anonymous function called")
	}()

	// Anonymous function
	func(value int) {
		fmt.Println("Your value is ", value)
	}(10)

	// Anonymous function returning value
	value := func(num1 int, num2 int) int {
		return num1 + num2
	}(1, 2)
	fmt.Println("Your addition is ", value)
}

func greet() {
	fmt.Println("Hello world")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(nums ...int) int {
	sum := 0
	for i := range nums {
		sum += i
	}
	return sum
}

func _2ReturnValue(num1 int, num2 int) (int, string) {
	return num1 + num2, "Hello bro"
}

func getArray() []int {
	return []int{1, 2, 3}
}
