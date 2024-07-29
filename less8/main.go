package main

import "fmt"

func main() {
	// slice1 := []int{1, 2, 3}
	// slice2 := []int{4, 5}
	// slice3 := append(slice1, slice2...) // добавление среза slice2 в срез slice1
	// fmt.Println(len(slice3), cap(slice3))
	// slice4 := append(slice2, slice1...) // добавление среза slice2 в срез slice1
	// fmt.Println(len(slice4), cap(slice4))

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:3:4]                 // длина 2, вместимость 3
	// fmt.Println(len(slice), cap(slice)) // 2 3
	// fmt.Println(slice)
	// fmt.Println("1:2:4")
	// slice2 :=arr[1:2:4]
	// fmt.Println(len(slice2), cap(slice2))
	// fmt.Println(slice2)

	// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7.
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr = ", arr)

	// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".
	var StringArr [5]string
	StringArr[0] = "a"
	StringArr[1] = "b"
	StringArr[2] = "c"
	StringArr[3] = "d"
	StringArr[4] = "e"
	fmt.Println("StringArr = ", StringArr)

	// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false.
	var boolArr = [4]bool{true, false, true, false}
	fmt.Println("boolArr", boolArr)

	// Создайте массив из 10 целых чисел без инициализации и выведите его на экран.
	var numbers [10]int
	fmt.Println(numbers)

	// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3.
	var boolArr2 = [4]bool{true, false, true, false}
	fmt.Println(boolArr2[1])
	fmt.Println(boolArr2[3])

	

	/*
			Array19. Дан целочисленный массив A размера 10. Вывести порядковый номер
		 последнего из тех его элементов AK, которые удовлетворяют двойному
		 неравенству A1 < AK < A10. Если таких элементов нет, то вывести 0.
	*/
	fmt.Println("Array 19  ")
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	idx := 0
	for i, v := range a {
		if a[0] < v && v < a[9] {
			idx = i
		}
	}
	fmt.Println(idx)

}
