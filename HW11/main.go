package main

import (
	"fmt"
	// "strings"
)

func main() {
	// 1) Конкатенация строк
	str1 := "Hi!"
	str2 := "Sher"
	fmt.Println(Concaten(str1, str2))

	// 2) Длина строки
	str_2 := "Sherzod"
	fmt.Println(returnLen(str_2))

	// 3) Проверка подстроки
	str_3 := "Hello World!"
	result_3 := "World"

	if containsSubstring(str_3, result_3) {
		fmt.Println("Подстрока найдена")
	} else {
		fmt.Println("Подстрока не найдена")
	}

	// 4) Поиск подстроки
	fmt.Println("4) Поиск подстроки")
	str_4 := "Hello, world!"
	substr_4 := "world"

	index_4 := findSubstringIndex(str_4, substr_4)

	if index_4 != -1 {
		fmt.Printf("Подстрока '%s' найдена в позиции %d\n", substr_4, index_4)
	} else {
		fmt.Println("Подстрока не найдена")
	}
	fmt.Println("index_4: ", index_4)

	// 5) Замена подстроки.
	fmt.Println("5) Замена подстроки.")
	text_5 := "Hello, world! Hello, Go! Hello Sher!"
	fmt.Println("text_5: ", text_5)
	find_5 := "Hello"
	new_5 := "Hi"
	result_5 := replaceAllSupStr(text_5, find_5, new_5)
	fmt.Println("result_5", result_5)

	// 6) Удаление пробелов.
	fmt.Println("6) Удаление пробелов.")
	str_6 := " Hello World! "
	fmt.Println("str_6: ", str_6)
	result_6 := trimStr(str_6)
	fmt.Println("result_6: ", result_6)

	// 7) Преобразование регистра.
	str_7 := "Hello Sher"
	upper, lower := convertToTo(str_7)
	fmt.Println("Верхний регистр:", upper)
	fmt.Println("Нижний регистр:", lower)

	// 8) Повторение строки.
	fmt.Println("8) Повторение строки.")
	str_8 := "Hello, World! Hello, Sher"
	fmt.Println("str_8: ", str_8)
	count_8 := 4

	result_8 := repeatStr(str_8, count_8)
	fmt.Println("result_8: ", result_8)

	// 9) Разбиение строки.
	fmt.Println("9) Разбиение строки.")
	str_9 := "apple,banana,orange"
	fmt.Println("str_9: ", str_9)
	separator_9 := ","

	str_9_2 := "Sher Sher Sher Sher"
	fmt.Println("str_9_2: ", str_9_2)
	separator_9_2 := " "

	result_9 := splitStr(str_9, separator_9)
	fmt.Println("result_9: ", result_9)
	result_9_2 := splitStr(str_9_2, separator_9_2)
	fmt.Println("result_9_2: ", result_9_2)

	//  10)Сравнение строк.
	fmt.Println(" 10)Сравнение строк.")
	text_10UP := "A"
	text_10LO := "a"

	result_10 := compareEqualFold(text_10UP, text_10LO)
	fmt.Println("result_10", result_10)

	// 11)Поиск и замена первой подстроки.
	fmt.Println("11)Поиск и замена первой подстроки.")
	text_11 := "Hello, World! Hello, Go!"
	fmt.Println("text: ", text_11)
	find_11 := "Hello"
	new_11 := "Hi"

	result_11 := replaceFirst(text_11, find_11, new_11)
	fmt.Println("result_11: ", result_11)

	// 12) Инверсия строки.
	fmt.Println("12) Инверсия строки.")
	str_12 := "Hello World!"
	reversed := reverseStrRune(str_12)
	fmt.Println("str_12: ", str_12)
	fmt.Println("reversed: ", reversed)

	// 13) Подсчет символов.
	fmt.Println("13) Подсчет символов.")
	str_13 := "hello, world!"
	fmt.Println("str_13: ", str_13)
	charCount := countChars(str_13)
	fmt.Println(charCount)

	// 14) Удаление символов.
	fmt.Println(" 14) Удаление символов.")
	fmt.Println(removeChar("Hello, World!", 'o'))

	// 15) Подсчет слов.
	fmt.Println("15) Подсчет слов.")
	str_15 := "Это пример строки для подсчета слов"
	fmt.Println(str_15)
	wordCount := countWorks(str_15)
	fmt.Println("wordCound: ", wordCount)

	// 16) Проверка префикса и суффикса.
	str_16 := "hello, world!"
	prifix := "hello"
	suffix := "world!"

	fmt.Println(hasPrefix(str_16, prifix)) // true
	fmt.Println(hasSuffix(str_16, suffix))

	// 17) Удаление дубликатов символов
	fmt.Println(" 17) Удаление дубликатов символов.")
	str_17 := "hello, world!"
	fmt.Println(removerDupl(str_17))

	// 18) Форматирование строки.
	fmt.Println("18) Форматирование строки.")
	str_18 := "Это строка с <тегами> и &символами&"
	formatted := formatHTMJl(str_18)

	fmt.Println("str_18: ", str_18)
	fmt.Println("formatted: ", formatted)

	// 19) Псевдонимы строк.
	fmt.Println("19) Псевдонимы строк.")
	str_19 := "Это пример строки с пробелами и спецсимволами!"
	slug_19 := createSlug(str_19)
	fmt.Println("str_19: ", str_19)
	fmt.Println("slug_19: ", slug_19)

	// 21) Обратный порядок слов.
	fmt.Println(" 21) Обратный порядок слов.")
	str_21 := "Это пример строки с обратным порядком слов"
	reversed_21 := reverseWords(str_21)
	fmt.Println("str_21: ", str_21)
	fmt.Println("reversed_21", reversed_21)

	// 22) Подсчет уникальных слов.
	fmt.Println("Отмена = 22) Подсчет уникальных слов. ")

	// 24) Перемешивание слов.
	fmt.Println("24) Перемешивание слов.")
	str_24 := "Это пример строки с перемешиваемыми словами"
	fmt.Println("str_24: ", str_24)
	shufflesStr := shuffleWords(str_24)
	fmt.Println("shufflesStr:", shufflesStr)

	// 25) Сортировка слов по длине.
	fmt.Println("25) Сортировка слов по длине.")
	str_25 := "Это пример строки сортировка слов по длине."
	sortStr_25 := sordWordsLength(str_25)
	fmt.Println("str_25: ", str_25)
	fmt.Println("sortStr_25: ", sortStr_25)

	// 29) Подсчет слов и символов.
	fmt.Println("29) Подсчет слов и символов.")
	str_29 := "Это пример строки подсчет слов и символов."
	fmt.Println("str_29: ")
	// fmt.Println(countWordsCharactes(str_29))
	words_29, char_29 := countWordsCharactes(str_29)
	fmt.Printf("Количество слов: %d и символов: %d", words_29, char_29)
}
