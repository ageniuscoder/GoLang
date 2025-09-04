package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "user@gmail.com",
	}
	fmt.Println(person)

	jsonData, err := json.Marshal(person) //marshal the struct to JSON
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))
	var person1 Person
	json.Unmarshal(jsonData, &person1) //unmarshal the JSON back to struct
	fmt.Println("Unmarshalled struct:", person1)
}
