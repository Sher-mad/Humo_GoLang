package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string  `json:"name`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}
type Person2 struct {
	Name      string `json:"name`
	Age       int    `json:"age"`
	IsStudent bool   `json:"IsStudent"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func main() {
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john.doe@example.com",
		Address: Address{
			Street: "123 Main St",
			City:   "Dushanbe",
		},
	}
	data, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile("example.json", os.O_CREATE|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	file.Write(data)

}

/*
data, err = json.MarshalIndent(book, "", " ")
if err != nil {
 fmt.Println(err)
}

file2, err := os.OpenFile("example2.json", os.O_CREATE|os.O_RDWR|os.O_CREATE, 0777)
if err != nil {
 panic(err)
}
file2.Write(data)

file, err = os.OpenFile("example3.json", os.O_CREATE|os.O_RDWR|os.O_CREATE, 0777)
if err != nil {
 panic(err)
}

*/
