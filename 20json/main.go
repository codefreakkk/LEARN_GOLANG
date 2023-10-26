package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"` // aliases
	Price    int
	Platform string
	Password string `json:"-"` // "-" by using this Password field will not be considered
	Tags     []string
}

func main() {
	JsonEncode()
	JsonDecode()
}

func JsonEncode() {
	courses := []course{{"Go lang", 1000, "lco", "harshsaid", []string{"go", "prog"}}}

	// Marshal() encodes into json
	finalCourses1, _ := json.Marshal(courses)
	fmt.Printf("%s", finalCourses1)

	finalCourses2, _ := json.MarshalIndent(courses, "", "\t")
	fmt.Printf("%s", finalCourses2)

}

// Consume JSON data received from WEB
func JsonDecode() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "Go lang",
		"Price": 1000,
		"Platform": "lco",
		"Tags": ["go","prog"]
	}
	`)

	var lcoCourse course
	isValid := json.Valid(jsonDataFromWeb)

	if !isValid {
		fmt.Println("JSON not valid")
		return
	}

	// decode json data and store in struct
	json.Unmarshal(jsonDataFromWeb, &lcoCourse) // decoding json data and storing in struct
	fmt.Printf("%#v\n", lcoCourse)              // printing whole json data which was stored in struct
	fmt.Println("Name of course is ", lcoCourse.Name)

	// decode json data and store in map
	var data map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &data)
	for k, v := range data {
		fmt.Printf("Key is -> %v, Value is -> %v\n", k, v)
	}
}
