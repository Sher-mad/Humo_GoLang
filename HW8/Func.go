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
	// multiplys := []int{}
	if len(num) == 0 {
		return zero
	}
	for i, v := range num {
		// multiply = append(multiplys, v*multiply)
		num[i] = v * multiply
	}
	return num
}

// 8. Найти все индексы заданного числа в массиве.
func fintInidics(arr []int, s int) []int {

	zero := []int{0}
	if len(arr) == 0 {
		return zero
	}

	result := []int{}
	for i, v := range arr {
		if v == s {
			result = append(result, i)

		}
	}
	return result
}

// 9. Создать копию массива.
func copyArray9(arr []int) []int {
	dst := make([]int, len(arr))
	copy(dst, arr)
	return dst
}

// 10. Объединить два массива.
// Задача 10 решено без функции

// 11. Поменять местами максимальный и минимальный элементы массива.
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func swapMinMax(arr []int) {
	if len(arr) == 0 {
		return // Обработка пустого массива
	}

	maxIndex, minIndex := 0, 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	swap(arr, maxIndex, minIndex)
}

//  12.	Проверить, является ли массив палиндромом.
func Palindrom(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

// 13. Найти второе наибольшее число в массиве.
func findSecondMaxArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max_1, max_2 := arr[0], arr[0]
	for _, v := range arr {
		if v > max_1 {
			max_2 = max_1
			max_1 = v
		} else if v > max_2 && v < max_1 {
			max_2 = v
		}
	}
	return max_2
}

// 14.	Перевернуть массив.
func PervernutArr(arr []int) []int {
	num := []int{}
	for i := 0; i < len(arr); i++ {
		num = append(num, arr[len(arr)-i-1])
	}
	return num
}

// // 15.	Удалить дубликаты из массива.
// func deleteElementArrDuble(arr []int, dubl int) []int {
// 	result := []int{}
// 	for i, v := range arr {
// 		// found = false
// 		if v == dubl {
// 			result[i] = arr[i]
// 		}

// 	}
// 	return result

// }
// 15 ДЗ не решено до конца


// 16.	Переместить все нули в конце массива, сохраняя порядок ненулевых элементов.
func ZeroFins(arr []int)


// 17.	Найти пересечение двух массивов.
// 18.	Проверить, является ли массив подмножеством другого массива.
// 19.	Объединить два отсортированных массива в один отсортированный.
// 20.	Найти длину самого длинного подмассива, в котором все элементы различны.
// 21.	Найти все подмассивы, сумма которых равна заданному числу.
// 22.	Найти пару элементов в массиве, сумма которых равна заданному числу.
// 23.	Найти наименьший положительный элемент, отсутствующий в массиве.
// 24.	Найти максимальную сумму подмассива с условием, что подмассив не должен содержать более двух различных элементов.
// 25.	Найти максимальную длину подмассива, сумма элементов которого равна заданному числу.
// 26.	Найти максимальное произведение трех чисел в массиве.
// 27.	Найти подмассив с максимальной суммой.
// 28.	Переместить все отрицательные числа в начало массива, сохраняя порядок остальных чисел.
// 29.	Найти подмассив с наибольшей длиной, сумма элементов которого равна нулю.
// 30.	Найти наибольший общий делитель всех элементов массива.
