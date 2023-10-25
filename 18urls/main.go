package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://lco.dev:3000/learn?cname=golang&paymentid=sjdkafhksajdf"

	result, _ := url.Parse(rawURL)

	fmt.Println(result)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	params := result.Query() // return's map
	for _, value := range params {
		fmt.Println("Params Value", value)
	}
}
