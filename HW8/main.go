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
	numbersArr2 := []int{3, 5, 9, 6, 7, 2}
	resultMinArrMax := findMinArr(numbersArr2)
	fmt.Println("Arr = ", numbersArr2)
	fmt.Println("Min arr = ", resultMinArrMax)

	// 3. Подсчитать количество положительных чисел в массиве.
	fmt.Println("3. Подсчитать количество положительных чисел в массиве.")
	numbersArr3 := []int{3, -5, 9, -6, 7, -2}
	resultFindArrCheck := findCheck(numbersArr3)
	fmt.Println("Arr = ", numbersArr3)
	fmt.Println("Check = ", resultFindArrCheck, "столько положттелных чисел")

	// 4. Найти сумму всех элементов массива.
	fmt.Println("4. Найти сумму всех элементов массива.")
	numbersArr4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultSumArr := AmountArr(numbersArr4)
	fmt.Println("Array = ", numbersArr4)
	fmt.Println("Sum Array = ", resultSumArr)

	// 5. Найти среднее значение всех элементов массива.
	fmt.Println("5. Найти среднее значение всех элементов массива.")
	numbersArr5 := []float64{1.5, 3.5, 5.5, 7.5, 9.5, 2.5, 6.5}
	resultArr5 := findAverage(numbersArr5)
	fmt.Println("Array 5 = ", numbersArr5)
	fmt.Printf("Averge = %.3f", resultArr5)
	fmt.Println("Averge = ", resultArr5)

	// 6. Удалить все вхождения заданного числа из массива.
	fmt.Println("6. Удалить все вхождения заданного числа из массива.")
	numbersArr6 := []int{2, 1, 2, 7, 3, 5, 2, 4, 2}
	fmt.Println("Array 6 = ", numbersArr6)
	dele := 2
	resultArr6, resultArr7 := deleteElementArr(numbersArr6, dele)
	// fmt.Println("Array 6 = ", numbersArr6)
	fmt.Println("Из масифа удалени элементи ", dele, "\n After", resultArr6)
	fmt.Println("Из масифа удалени элементи ", dele, "\n After", resultArr7)

}
