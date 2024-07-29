package main

import "fmt"

func main() {
	//1. Найти максимальный элемент в массиве.
	fmt.Println("1. Найти максимальный элемент в массиве.")
	numbersMAx := []int{3, 5, 9, 6, 7, 2}
	resultFindArrMax := findMAxArr(numbersMAx)
	fmt.Println(" Arr = ", numbersMAx)
	fmt.Println("Max Arr = ", resultFindArrMax)

	// 2. Найти минимальный элемент в массиве.
	fmt.Println("2. Найти минимальный элемент в массиве.")
	numbersMin := []int{3, 5, 9, 6, 7, 2}
	resultMinArrMax := findMinArr(numbersMin)
	fmt.Println("Arr = ", numbersMin)
	fmt.Println("Min arr = ", resultMinArrMax)

}
