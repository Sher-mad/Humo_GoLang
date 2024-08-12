package main

import "fmt"

func main() {
	//
	fmt.Println("1. Книга и Автор")
	book1 := Book{
		title:  "Sumerki",
		auther: "Sherzod",
	}
	book2 := Book{title: "Narni", auther: "Ali"}

	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)

	for i := 0; i < len(library.books); i++ {
		fmt.Println(Book.GetDetails(library.books[i]))
	}

	// 2. Оценки студента
	fmt.Println("2. Оценки студента")
	studend1 := Studend{}
	studend1.AddGrade(50)
	studend1.AddGrade(70)
	studend1.AddGrade(80)
	studend1.AddGrade(90)
	average := studend1.AverageGrade()
	fmt.Printf("Средний балл студента: %.2f \n ", average)

	// 3. Круг и Площадь
	fmt.Println("3. Круг и Площадь")
	radius := Circle{radius: 5}

	area := radius.Area()
	circum := radius.Circumference()

	fmt.Printf("Area = %.4f\n ", area)
	fmt.Printf("Circle = %.4f\n", circum)

	// 4. Контейнер для товаров
	fmt.Println("4. Контейнер для товаров")
	item1 := Item{"Яблоко", 1.5}
	item2 := Item{"Банан", 2.0}
	item3 := Item{"Апельсин", 1.8}

	inventory := Inventory{}
	inventory.AddItem(item1)
	inventory.AddItem(item2)
	inventory.AddItem(item3)

	fmt.Println("Array Item = ", inventory)
	total := inventory.Sum()
	fmt.Printf("Общая стоимость товаров: %.2f\n", total)

	// 5. Обработка транзакций
	fmt.Println("5. Обработка транзакций")
	account := Account{}

	account.AddTransaction(100.0, "Пополнение счета")
	account.AddTransaction(-50.0, "Снятие наличных")
	// fmt.Println("Account: ", account.transaction)
	for _, transaction := range account.transaction {
		transaction.Process()
		fmt.Println(transaction)
	}

	// 6. Управление задачами
	fmt.Println("6. Управление задачами")
	tasksList := TaskList{}
	tasksList.AddTask("Купить молоко")
	tasksList.AddTask("Вынести мусор")
	tasksList.AddTask("Позвонить другу")

	tasksList.tasks[0].copmleted = true
	fmt.Printf("Количество завершенных задач: %d\n", tasksList.CompletedTasks())

	// 7. Работа с температурой
	fmt.Println("7. Работа с температурой")

	temp := Temperature{25}
	fmt.Printf("%.2f градусов Цельсия равняется %.2f градусам Фаренгейта и %.2f Кельвинам.\n", temp.celsius, temp.ToFahrenheit(), temp.ToKelvin())

	// 8. Управление списком студентов
	fmt.Println("8. Управление списком студентов")
	classroom_8 := Classrom{}
	classroom_8.AddStudent("Sher", 26)
	classroom_8.AddStudent("Abu", 22)
	classroom_8.AddStudent("Ali", 24)
	fmt.Printf("Total Quantity: %.2f\n", classroom_8.AverageAge())

	// 9. Склад товаров
	fmt.Println("9. Склад товаров")
	wareHouse := Warehouse{}
	wareHouse.AddProduct("iPhone 15 max", 100)
	wareHouse.AddProduct("Samsund S24 ultra", 120)
	wareHouse.AddProduct("Mi 14 Ultra ", 80)
	fmt.Printf("Total Quantity: %d\n", wareHouse.TotalQuantity())

	// 10. Заметки и метки
	fmt.Println("10. Заметки и метки")
	noteBook_10 := Notebook{}
	noteBook_10.AddNote("Сделать покупки книг", []string{"личный", "задачи"})
	noteBook_10.AddNote("Прочитать книгу", []string{"хобби", "развитие"})
	noteBook_10.AddNote("Позвонить другу", []string{"личные", "связь"})
	noteBook_10.AddNote("Сделать покупки ноутбуков", []string{"личный", "задачи"})
	noteBook_10.AddNote("Сделать покупки товаров", []string{"личный", "задачи"})

	noteWithTag := noteBook_10.NotesWithTag("личный")
	fmt.Println("Заметки с меткой 'личные':")
	for _, note := range noteWithTag {
		fmt.Println(note.content)
	}

	// 11. Управление зарплатами
	fmt.Println("11. Управление зарплатами")
	employees_11 := Payroll{}
	employees_11.AddEmployee("Sher", 6000.0)
	employees_11.AddEmployee("Ali", 5000.0)
	employees_11.AddEmployee("Abu", 5500.0)
	for _, emp := range employees_11.employees {
		fmt.Println(emp)
	}
	averPay_11 := employees_11.AverageSalary()
	fmt.Printf("Средняя зарплата: %.2f\n", averPay_11)

	// 12. Разделение бюджета
	fmt.Println("12. Разделение бюджета")
	budget_12 := Budget{amount: 6000}
	if budget_12.Allocate(5000) {
		fmt.Println("Allocation successful")
	} else {
		fmt.Println("Allocation failed")
	}
	fmt.Printf("Remaining Budget: %.2f\n", budget_12.Remaining())

	budgets_12 := Budget{amount: 1000}
	fmt.Println("Изначальный бюджет:", budgets_12.Remaining())

	allocate1 := budgets_12.Allocate(800)
	fmt.Println("Выделено 800:", allocate1)
	fmt.Println("Остаток: ", budgets_12.Remaining())

	allocate2 := budgets_12.Allocate(300)
	fmt.Println("Выделено 300: ", allocate2)
	fmt.Println("Остаток: ", budgets_12.Remaining())

	// 13. Отслеживание заказов
	fmt.Println("13. Отслеживание заказов")
	orders := OrderManager{}
	orders.AddOrder(1, 20.5)
	orders.AddOrder(2, 22.5)
	orders.AddOrder(3, 24.5)
	orders.AddOrder(4, 26.5)
	fmt.Println("Покупки:")
	for _, ord := range orders.orders {
		fmt.Printf("%d: %.2f$\n", ord.id, ord.totalAmount)
	}
	fmt.Printf("Сумма покупки: %.2f$\n", orders.TotalOrdersAmount())

	// 14. Управление аккаунтами
	fmt.Println("14. Управление аккаунтами")
	account14 := Account14{balance: 1000}
	account14.Deposit(500)
	if account14.Withdraw(1600) {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Withdrawal failed")
	}
	fmt.Printf("Current Balance: %.2f\n", account14.GetBalance())

	// 15. Операции над строками
	fmt.Println("15. Операции над строками")
	text_15 := Text{}
	text_15.AddText("Hello, Sher.")
	text_15.AddText("Welcome to Go Programming.")
	fmt.Println("text_15: ", text_15.content)
	text_15.ReplaceWord("Sher", "World!")
	fmt.Printf("Text Length: %d\n", text_15.Length())
	fmt.Println("Text Content: ", text_15.content)
}
