package main

import (
	"fmt"
	"os"
)

func main() {
	// Открытие или создание файла для записи
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error Create file:", err)
	// 	return
	// }
	// defer file.Close()

	// data, _ := io.ReadAll(file)
	// fmt.Println(data)

	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error Create file:", err)
		return
	}
	defer file.Close()
	fmt.Println("file opened successfully")

}
