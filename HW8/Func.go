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
