package main

import "fmt"

func main() {
	// 1. Создание и вывод map
	fmt.Println("1. Создание и вывод map")
	m_1 := make(map[string]int)
	m_1["apple"] = len("apple")
	m_1["banana"] = len("banana")
	printMap_1(m_1)

	// 2. Проверка наличия ключа
	fmt.Println("2. Проверка наличия ключа")
	m2 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println("apple exists:", keyExists_2(m2, "apple"))
	fmt.Println("grape exists:", keyExists_2(m2, "grape"))

	// 3. Обновление значения по ключу
	fmt.Println("3. Обновление значения по ключу")
	m3 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(m3)
	updateValue(m3, "banana", 3)
	fmt.Println(m3)
}
