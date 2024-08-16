package main

import (
	"fmt"
	"sort"
	"strings"
)

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

// 4. Удаление элемента из map
func deledeElement(m map[string]int, key string) {
	delete(m, key)
}

// 5. Подсчет частоты слов
func wordFrequency(text string) map[string]int {
	frequency := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

// 6. Сортировка ключей в map
func sortedKey(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)

	}
	sort.Strings(keys)
	return keys
}

// 7. Проверка пустого map
func isMapEmpty(m map[string]int) bool {
	return len(m) == 0
}

// 8. Инвертирование ключей и значений
func invertMap(m map[string]int) map[int]string {
	inverted := make(map[int]string)
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}

// 9. Проверка дубликатов в map
func hasDublicates(m map[string]int) bool {
	valueSet := make(map[int]struct{})
	for _, value := range m {
		if _, exists := valueSet[value]; exists {
			return true
		}
		valueSet[value] = struct{}{}
	}
	return false
}
