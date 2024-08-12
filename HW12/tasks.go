package main

import "fmt"

// 1. Создание и вывод map
func printMap_1(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// 2. Проверка наличия ключа
func keyExists_2(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// 3. Обновление значения по ключу
func updateValue(m map[string]int, key string, newValue int) {
	m[key] = newValue
}
