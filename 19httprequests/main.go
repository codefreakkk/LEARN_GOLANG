package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	GetRequest()
}

func GetRequest() {
	rawUrl := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(rawUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code", response.StatusCode)
	fmt.Println("Contenet length", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

	// anoter way to convert byte into string
	var responseString strings.Builder
	responseString.Write(content)
	fmt.Println(responseString.String())

}
