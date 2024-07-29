package main

import "fmt"

func main() {
	// // 1 Temperature Проверка температуры
	// fmt.Println("1 Temperature Проверка температуры")
	// temp1 := Temperature(20.5)
	// temp2 := Temperature(-8.5)
	// temp3 := Temperature(0.0)

	// fmt.Println(proverkaTemperature(temp1))
	// fmt.Println(proverkaTemperature(temp2))
	// fmt.Println(proverkaTemperature(temp3))

	// // 2 vozAge определает возрасть
	// var vozraAge vozrastAge
	// fmt.Print("\n Сколько вам лет? ")
	// fmt.Scan(&vozraAge)
	// fmt.Println(VozAge(vozraAge))

	// // 3 sk Проверка скорости.
	// fmt.Println("3 sk Проверка скорости")
	// speed1 := Speed(100)
	// speed2 := Speed(-20)
	// speed3 := Speed(130)
	// fmt.Println(ProverSpeed(speed1))
	// fmt.Println(ProverSpeed(speed2))
	// fmt.Println(ProverSpeed(speed3))

	// //4. Soce Проверка счета.
	// fmt.Println("4. Soce Проверка счета.")
	// cheksScore1 := Score(2)
	// cheksScore2 := Score(2)
	// cheksScore3 := Score(2)

	// chekScore(cheksScore1)
	// chekScore(cheksScore2)
	// chekScore(cheksScore3)

	// //5. BMI Классификация
	// fmt.Println("5. BMI	Классификация ")
	// var vesBMI BMI = 78.5
	// category := classifyBMI(vesBMI)
	// fmt.Println("Ваш BMI = ", vesBMI, ", category =", category)

	// 6. Возведение в квадрат.
	fmt.Println("6. Возведение в квадрат.")
	// Создания переменой типа UnaryOperation и
	// присваиваем ей функцию kvatrat
	up := UnaryOperation(kvatrat)
	//вызов функции через up
	result := up(4)
	fmt.Println("6) result =", result)

	//7. Удвоение числа
	fmt.Println("7. Удвоение числа")
	up2 := UnaryOperation2(udvoenie)
	result2 := up2(4)
	fmt.Println("7) result2 =", result2)

	// 8. Проверка четности.
	fmt.Println("8. Проверка четности.")
	ostatok := checkOstatok(CheckOstatok)
	result3 := ostatok(20)
	fmt.Println("8) result3", result3)

	// 9. Проверка положительности.
	fmt.Println("9. Проверка положительности.")
	checkPoloj := CheckPoloj(15)
	fmt.Println("checkPoloj = ", checkPoloj)

	// 10. Комбинирование функций. Не смог

	// 11.	Обратный отсчет
	fmt.Println("11. Обратный отсчет")
	startCount := Count(5)
	countObratni(startCount)

	// 14. Оценка успеваемости.
	fmt.Println("14. Оценка успеваемости")
	grade := Grade(49)
	Grade1(grade)

	// 15. Оценка состояния здоровья
	fmt.Println("15. Оценка состояния здоровья")
	vashBMI := BMI(25.4)
	vashBP := BloodPressure(84)
	resultHealth := proverkaHealth(vashBMI, vashBP)
	fmt.Println("", resultHealth)

	// 16. Описание продукта.
	fmt.Println("16. Описание продукта.")
	product16X := Product16{Name: "Apple", Price: 15000}
	result16 := product16(product16X)
	fmt.Println(result16)

	// 17.	Вывод информации о человеке.
	fmt.Println("17. Вывод информации o человеке.")
	person17s := Person17{FirstName: "Sherzod", LastName: "Qurbonov", Age: 26}
	// person(persons)
	result17 := person17(person17s)
	fmt.Println(result17)

	// 18. Сравнение продуктов.
	fmt.Println("18. Сравнение продуктов.")
	product18_1 := Product18{Name: "Apple", Price: 15}
	product18_2 := Product18{Name: "Carrot", Price: 5}
	// product18(product18_1, product18_2)
	result18 := product18(product18_1, product18_2)
	fmt.Println("Дорогой продуктом стал = ", result18)

	// 21. Обновление цены продукта
	fmt.Println("21. Обновление цены продукта")
	product21x := Product21{Name: "iPhone 14", Price: 1355.99}
	fmt.Println("измениния цены = ", product21x.Name, " = ", product21x.Price)
	product21(&product21x, 1299.99)
	fmt.Println("Цена изменено = ", product21x.Price)

	// 22. Увеличение возраста.
	fmt.Println("22. Увеличение возраста.")
	person22x := Person22{Name: "Sherzod", Age: 26}
	fmt.Println("Возраст = ", person22x.Age)
	person22(&person22x)
	fmt.Println("Возраст = ", person22x.Age)

	// 23. Обновление информации о продукте.
	fmt.Println("23. Обновление информации о продукте.")
	product23x := Product23{Name: "'iPHone 14' ", Price: 1495.99}
	fmt.Println("Before = ", product23x)
	newName23 := "'Samsung'"
	newPrice23 := 1699.99
	product23(&product23x, newName23, newPrice23)
	fmt.Println("After = ", product23x)

	//24. Увеличение количества предметов.
	fmt.Println("24. Увеличение количества предметов.")
	Item24x := Item24{NAme: "ЩИТ", Quantity: 8}
	fmt.Println("Before = ", Item24x)
	amount := 2
	item24(&Item24x, amount)
	fmt.Println("After = ", Item24x)

	// 26. Адрес человека.
	address24x := Person24{Name: "Sherzod", Address: Address24{Street: "Tajikistan", City: "Dushanbe"}}
	fmt.Println(address24(address24x))

	// 27. Описание компании
	fmt.Println("27. Описание компании")
	company27x := Company27{
		Name:     "SherTechnoli",
		Location: Address24{Street: "Айни 155", City: "Dushanbe"},
	}
	result27 := company27(company27x)
	fmt.Println(result27)

	// //30. Управление проектом.
	// fmt.Println("30. Управление проектом.")
	// project30x := Project30{
	// 	Name: "Delevely",
	// 	Meneger: Person30{
	// 		Name: "Sherzod",
	// 		Address: Address30{
	// 			Street: "Айни 155",
	// 			City:   "Dushabe",
	// 		},
	// 	},
	// }
	// result30 := project30(project30x)
}
