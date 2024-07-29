package main

import "fmt"

// import "fmt"

// /*
// 	1.	Проверка температуры
// 	Определите тип Temperature на основе float64.
// 	Напишите функцию, которая принимает температуру и возвращает сообщение о том,
// 	является ли она ниже, выше или равной нулю.

// */

// // Temperature представляет температуру в градусах
// type Temperature float64

// // func (t Temperature) string() string {
// // 	return fmt.Sprint("%.2fvградусов", t)
// //}

// func proverkaTemperature(temp Temperature) string {
// 	switch {
// 	case temp > 0:
// 		return fmt.Sprintf("%.1f выше нуля", temp)
// 	case temp < 0:
// 		return fmt.Sprintf("%.1f ниже нуля", temp)
// 	default:
// 		return fmt.Sprintf("%.1f равна нулю", temp)
// 	}
// }

// /*
// 2.Проверка возраста Определите тип Age на основе int.
// Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек подростком (возраст от 13 до 19 лет включительно) или нет.
// */
// // VozAge определает возрасть
// type vozrastAge int

// func VozAge(age vozrastAge) string {
// 	if age >= 18 {
// 		return "У вас имеется лиценция совершеннолетия"
// 	} else if age > 13 && age < 19 {
// 		return "Вам чуть чуть осталось чтоб получить лиценцию и совершеннолетия, Но вы евляетесь подростком"
// 	} else {
// 		return "Вам ещё рано на триторию совершеннолетных"
// 	}

// }

// /*
// 3. Проверка скорости.
// Определите тип Speed на основе float64.
// Напишите функцию, которая принимает скорость и возвращает сообщение о том,
// является ли она допустимой (от 0 до 120 км/ч включительно) или нет.
// */

// type Speed float64

// func (s Speed) String() string {
// 	return fmt.Sprintf("%.2f км/ч", s)
// }
// func boolProverkaSpeed(speed Speed) bool {
// 	return speed >= 0 && speed <= 120
// }
// func ProverSpeed(speed Speed) string {
// 	if boolProverkaSpeed(speed) {
// 		return fmt.Sprintf("%s - допустимая скорость", speed)
// 	} else {
// 		return fmt.Sprintf("%s - недопустимая скорость", speed)
// 	}
// }

// /*
// 4.	Проверка счета. Определите тип Score на основе int.
// Напишите функцию, которая принимает счет и возвращает сообщение о том,
// является ли он положительным, отрицательным или нулевым.
// */
// type Score int

// func chekScore(num Score) string {
// 	switch {
// 	case num > 0:
// 		return fmt.Sprintf("%d - явлается положительным. ", num)
// 	case num < 0:
// 		return fmt.Sprintf("%d  -явлается отрицательным. ", num)
// 	default:
// 		return fmt.Sprintf("%d - явлается нулевым ", num)
// 	}
// }

/*
5.	Классификация BMI
Определите тип BMI на основе float64.
Напишите функцию, которая принимает значение BMI и возвращает сообщение о том,
в какой категории оно находится: "Underweight",
"Normal weight", "Overweight", или "Obesity".
*/
// type BMI float64

// func classifyBMI(bmi BMI) string {
// 	switch {
// 	case bmi < 55.0:
// 		return "Underweight"
// 	case bmi >= 55.0 && bmi < 85.0:
// 		return "Normal weight"
// 	case bmi >= 85.0 && bmi < 110.0:
// 		return "Overweight"
// 	default:
// 		return "Obesity"
// 	}
// }

/*
6. Возведение в квадрат.
Определите тип функции UnaryOperation,
которая принимает int и возвращает int.
Напишите функцию для возведения числа в квадрат и используйте
тип UnaryOperation для вызова этой функции.
*/
// UnaryOperation - тип для функций
type UnaryOperation func(int) int

func kvatrat(x int) int {
	return x * x
}

/*
7.	Удвоение числа
Определите тип функции UnaryOperation,
которая принимает int и возвращает int.
Напишите функцию для удвоения числа и используйте
тип UnaryOperation для вызова этой функции.
*/
type UnaryOperation2 func(int) int

func udvoenie(x int) int {
	return x * 2
}

/*
8. Проверка четности. Определите тип функции Check,
которая принимает int и возвращает bool.
Напишите функцию для проверки четности числа и
используйте тип Check для вызова этой функции.
*/
type checkOstatok func(int) bool

func CheckOstatok(a int) bool {
	switch {
	case a%2 == 0:
		return true
	default:
		return false
	}
}

/*
9.	Проверка положительности. Определите тип функции Check,
которая принимает int и возвращает bool. Напишите функцию для
проверки, является ли число положительным, и используйте
тип Check для вызова этой функции.
*/
func CheckPoloj(a int) bool {
	if a > 0 {
		return true
	} else {
		return false
	}

}

/*
10. Комбинирование функций.
Определите тип функции Operation, которая принимает два int и возвращает int.
Напишите функции для сложения, вычитания и умножения.
Напишите функцию, которая принимает два int и массив функций Operation,
и возвращает массив результатов применения этих функций.
*/
// Operation тип для функций
// type Operation func(int, int) int

// func Plus(a, s int) int {
// 	return a + s
// }
// func Minus(a, s int) int {
// 	return a - s
// }
// func Umnoj(a, s int) int {
// 	return a * s
// }

/*
11.	Обратный отсчет.Создайте псевдоним для типа int под названием Count.
Напишите функцию, которая принимает Count и выводит обратный отсчет до нуля.
*/
type Count int

func countObratni(count Count) {
	for count >= 0 {
		fmt.Println("Count = ", count)
		count--
	}
}

/*
14.	Оценка успеваемости.
Создайте псевдоним для типа int под названием Grade.
Напишите функцию, которая принимает Grade и возвращает сообщение о том,
является ли оценка удовлетворительной (50 и выше) или нет
*/
type Grade int

func Grade1(x Grade) {
	switch {
	case x >= 50:
		fmt.Println("Твоя оценка удовлетворительная = ", x)
	default:
		fmt.Println("Твоя не оценка удовлетворительная = ", x)
	}
}

/*
15.	Оценка состояния здоровья
Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure.
Напишите функцию, которая принимает BMI и BloodPressure,
и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy
*/
// псевдонимы nbgjd
type BMI float64
type BloodPressure float64

func proverkaHealth(bmi BMI, bp BloodPressure) string {
	switch {
	case bmi >= 19.5 && bmi < 23.5 && bp >= 70 && bp < 85:
		return "Healthy"
	case bmi >= 23.5 && bmi < 27.5 && bp >= 85 && bp < 100:
		return "At risk"
	case bmi >= 27.5 && bmi < 32.5 && bp >= 100 && bp < 140:
		return "Unhealthy"
	default:
		return "Вы труп"
	}
}

/*
16.	Описание продукта.
Создайте структуру Product с полями Name и Price.
Напишите функцию, которая принимает Product и выводит его описание.
*/

type Product16 struct {
	Name  string
	Price int
}

func product16(x Product16) Product16 {
	return x
}

/*
17.	Вывод информации о человеке.
Создайте структуру Person с полями FirstName, LastName и Age.
Напишите функцию, которая принимает Person и выводит его полное имя и возраст.
*/

type Person17 struct {
	FirstName string
	LastName  string
	Age       int
}

func person17(x Person17) Person17 {
	//return fmt.Sprintf("FirstName: %s, LastName: %s, Age: %s", x.FirstName, x.LastName, x.Age)
	return x
}

/*
18.	Сравнение продуктов.
Создайте структуру Product с полями Name и Price.
Напишите функцию, которая принимает два Product и возвращает
более дорогой продукт.
*/
type Product18 struct {
	Name  string
	Price int
}

func product18(x, c Product18) Product18 {
	// switch {
	// case x.Price > c.Price:
	// 	return fmt.Sprintf("%s = %s дороже %s", x.Name, x.Price, c.Name)
	// case x.Price < c.Price:
	// 	return fmt.Sprintf("%s =  %s дороже %s", c.Name, c.Price, x.Name)
	// default:
	// 	return fmt.Sprintf("Err")
	// }
	if x.Price > c.Price {
		return x
	} else {
		return c
	}
}

/*
21.	Обновление цены продукта
Создайте структуру Product с полями Name и Price. Напишите функцию,
которая принимает указатель на Product и обновляет его цену.
*/
type Product21 struct {
	Name  string
	Price float64
}

func product21(a *Product21, newPlice float64) {
	a.Price = newPlice
}

/*
22.	Увеличение возраста.
Создайте структуру Person с полями Name и Age. Напишите функцию,
которая принимает указатель на Person и увеличивает его возраст на 1.
*/
type Person22 struct {
	Name string
	Age  int
}

func person22(a *Person22) {
	a.Age++
}

/*
23. Обновление информации о продукте.
Создайте структуру Product с полями Name и Price. Напишите функцию,
которая принимает указатель на Product и обновляет его название и цену.
*/
type Product23 struct {
	Name  string
	Price float64
}

func product23(a *Product23, newName string, newPrice float64) {
	a.Name = newName
	a.Price = newPrice
}

/*
24. Увеличение количества предметов.
Создайте структуру Item с полями Name и Quantity. Напишите функцию,
которая принимает указатель на Item и увеличивает количество на заданное значение.
*/

type Item24 struct {
	NAme     string
	Quantity int
}

func item24(a *Item24, amount int) {
	a.Quantity += amount
}

/*
26. Адрес человека.
Создайте структуру Address с полями Street и City.
Создайте структуру Person с полями Name и Address. Напишите функцию,
которая принимает Person и выводит его имя и адрес.
*/
type Address24 struct {
	Street string
	City   string
}

type Person24 struct {
	Name    string
	Address Address24
}

func address24(s Person24) Person24 {
	return s
}

/*27.	Описание компании
Создайте структуру Company с полями Name и Location (структура Address).
Напишите функцию, которая принимает Company и выводит информацию о компании.
*/
type Company struct {
	Name     string
	Location Address24
}
// func category
