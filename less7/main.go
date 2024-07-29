package main

import (
	"fmt"
)

type Age int // Создания именованный тип на основе int
type Percon struct {
	Name string
	Age  int
}
type vozrastAge int

func main() {
	var myAge Age = 25
	fmt.Println(myAge)

	//struct
	// Создания и  нициализируем структуру
	// p:=Person{Name:"Sher", Age: 30}
	// Анонимные стуктуры
	// person := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Alice",
	// 	Age:  30,
	// }
	// fmt.Println(person)

	// // func vozAge
	// var vozraAge vozrastAge
	// fmt.Print("\n Сколько вам лет? ")
	// fmt.Scan(&vozraAge)
	// fmt.Println(VozAge(vozraAge))

	//func Number
	// var numAge Number
	// fmt.Println("func Number")
	// fmt.Scan(&numAge)
	// fmt.Println(number(numAge))

	//func score
	var Score1 Score
	fmt.Println("func score")
	fmt.Scan(&Score1)
	fmt.Println(score(Score1))

}

// Определение возраста совершеннолетия
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
func VozAge(age vozrastAge) string {
	if age >= 18 {
		return "У вас имеется лиценция совершеннолетия"
	} else if age > 15 && age < 18 {
		return "Вам чуть чуть осталось чтоб получить лиценцию и совершеннолетия"
	} else {
		return "Вам ещё рано на триторию совершеннолетных"
	}

}

/*
Проверка на четность. Определите тип Number на основе int.
Напишите функцию, которая принимает число и возвращает сообщение о том,
является ли оно четным или нечетным
*/
type Number int

func number(age Number) string {
	if age%2 == 0 {
		return "Число явлается четным"
	} else {
		return " Число не четное"
	}
}

/*
Проверка допустимого диапазона. Определите тип Score на основе int.
Напишите функцию, которая принимает оценку и возвращает сообщение,
находится ли она в допустимом диапазоне от 0 до 100
*/
type Score int

func score(age Score) string {
	if age < 100 && age > 0 {
		return "Ваша оценка в лоступном диапазоне вы прошли"
	} else if age > 100 {
		return "Ваша оценка више диапзона"
	} else {
		return "Ошибка"
	}
}
