package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGet() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/5")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(fmt.Errorf("error %v", res.Status))
		return
	}

	// data, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading data", err)
	// }
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Invalid data type")
	}

	fmt.Println(todo)

}

func performPost() {
	todo := Todo{
		UserId:    87,
		Title:     "mangal is here",
		Completed: false,
	}
	data, err := json.Marshal(&todo)
	if err != nil {
		fmt.Println("error converting data")
	}
	jsonData := string(data)
	jsonReader := strings.NewReader(jsonData)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error creating data")
		return
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error reading res")
		return
	}
	fmt.Println(string(data))

}

func performUpdate() {
	todo := Todo{
		UserId:    9883,
		Title:     "mangal is here",
		Completed: false,
	}
	data, err := json.Marshal(&todo)
	if err != nil {
		fmt.Println("error converting data")
	}
	jsonData := string(data)
	jsonReader := strings.NewReader(jsonData)
	const url = "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, url, jsonReader)

	if err != nil {
		fmt.Println("Error creating Put req")
		return
	}

	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending Request")
		return
	}
	defer res.Body.Close()

	data, _ = io.ReadAll(res.Body)
	fmt.Println(string(data))
}

func performDelete() {
	const url = "https://jsonplaceholder.typicode.com/todos/5"
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Errro sendind req")
		return
	}
	res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
	fmt.Println(res.Status)

}
func main() {
	//performGet()
	//performPost()

	//performUpdate()

	performDelete()

}
