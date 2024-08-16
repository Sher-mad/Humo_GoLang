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

	// 4. Удаление элемента из map
	fmt.Println("4. Удаление элемента из map")
	m4 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orenge": 3,
	}

	fmt.Println("До удаления: ", m4)
	deledeElement(m4, "apple")
	fmt.Println("После удаления: ", m4)

	// 5. Подсчет частоты слов
	fmt.Println("5. Подсчет частоты слов")
	text5 := "Это пример текста это пример"
	freqMap := wordFrequency(text5)

	for word, count := range freqMap {
		fmt.Printf("%s: %d\n", word, count)
	}

	// 6. Сортировка ключей в map
	fmt.Println("6. Сортировка ключей в map")
	m6 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
	sorted := sortedKey(m6)
	fmt.Println("Отсортировка ключи: ", sorted)

	// 7. Проверка пустого map
	fmt.Println("7. Проверка пустого map")
	m7 := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	emptyMap7 := map[string]int{}
	fmt.Println("M7 пустой? = ", isMapEmpty(m7))
	fmt.Println("emptyMap пустой? = -", isMapEmpty(emptyMap7))

	// 8. Инвертирование ключей и значений
	fmt.Println("8. Инвертирование ключей и значений")
	m8 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
		"Что-то": 4,
		"Да-Ну":  1,
	}
	fmt.Println("Before: ", m8)

	invertedMap := invertMap(m8)
	fmt.Println("After: ", invertedMap)

	// 9. Проверка дубликатов в map
	fmt.Println("9. Проверка дубликатов в map")
	/*
		Проверка дубликатов в map
		Описание: Напишите функцию, которая проверяет, есть
		ли дубликаты значений в map.
	*/

	m9 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
		"Что-то": 4,
		"Да-Ну":  2,
	}
	fmt.Println("Есть ли дубликаты в myMap? = ", hasDublicates(m9))

	anotherMap9 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
	fmt.Println("Есть ли дубликаты в anotherMap? = ", hasDublicates(anotherMap9))
}
