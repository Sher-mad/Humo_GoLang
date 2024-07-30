package main

/*
1.	Найти максимальный элемент в массиве.
*/

func findMAxArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

/*
2. Найти минимальный элемент в массиве.
*/
func findMinArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, num := range arr {
		if num < max {
			max = num
		}
	}
	return max
}

/*
3.	Подсчитать количество положительных чисел в массиве.
*/
func findCheck(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	check := 0
	for _, chec := range arr {
		if chec > 0 {
			check++
		}

	}
	return check

}

/*
4. Найти сумму всех элементов массива.
*/
func AmountArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	amountMax := 0
	for _, num := range arr {
		amountMax += num
	}
	return amountMax
}

/*
5. Найти среднее значение всех элементов массива.
*/
func findAverage(average []float64) float64 {
	if len(average) == 0 {
		return 0.0
	}
	var sum float64
	for _, v := range average {
		sum += v
	}
	return sum / float64(len(average))
}

/*
6. Удалить все вхождения заданного числа из массива.
*/
func deleteElementArr(arr []int, delete int) ([]int, []int) {
	if len(arr) == 0 {
		return arr, arr
	}
	slow := 0
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] != delete {
	// 		arr[slow] = arr[i]
	// 		slow++
	// 	}
	// }
	// return arr[:slow]

	for i, v := range arr {
		if v != delete {
			arr[slow] = arr[i]
			slow++
		}
	}
	return arr[:slow], arr

}

// 7. Умножить все элементы массива на заданное число.

func multiplyArray(num []int, multiply int) []int {
	zero := []int{0}
	if len(num) == 0 {
		return zero
	}
	return num
	
}
