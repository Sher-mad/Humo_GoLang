package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Book struct {
	title  string
	auther string
}
type Library struct {
	books []Book
}

// 1. Книга и Автор

func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Auther: %s", b.title, b.auther)
}
func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

// 2. Оценки студента
type Studend struct {
	grades []int
}

func (s *Studend) AddGrade(grade int) {
	s.grades = append(s.grades, grade)
}
func (s Studend) AverageGrade() float64 {
	if len(s.grades) == 0 {
		return 0
	}
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum) / float64(len(s.grades))
}

// 3. Круг и Площадь
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

// 4. Контейнер для товаров
type Item struct {
	Name  string
	Price float64
}
type Inventory struct {
	item []Item
}

func (i *Inventory) AddItem(item Item) {
	i.item = append(i.item, item)
}

func (i Inventory) Sum() float64 {
	var sum float64
	for _, item := range i.item {
		sum += item.Price
	}
	return sum
}

// 5. Обработка транзакций
type Transaction struct {
	amount      float64
	description string
}

func (t Transaction) Process() string {
	return fmt.Sprintf("Транзакция: %.2f  %s\n", t.amount, t.description)
}

type Account struct {
	transaction []Transaction
}

func (a *Account) AddTransaction(amount float64, description string) {
	newTransaction := Transaction{amount, description}
	a.transaction = append(a.transaction, newTransaction)
}

// 6. Управление задачами
type Task struct {
	title     string
	copmleted bool
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) AddTask(title string) {
	newTask := Task{title, false}
	t.tasks = append(t.tasks, newTask)
}
func (t TaskList) CompletedTasks() int {
	var completedCount int
	for _, task := range t.tasks {
		if task.copmleted {
			completedCount++
		}

	}
	return completedCount
}

// 7. Работа с температурой
type Temperature struct {
	celsius float64
}

// return t.celsius*9/5 + 32
func (t Temperature) ToFahrenheit() float64 {
	return t.celsius*9/5 + 32
}

func (t Temperature) ToKelvin() float64 {
	return t.celsius + 273.15
}

// 8. Управление списком студентов
type Studend8 struct {
	Name string
	Age  int
}
type Classrom struct {
	classroom []Studend8
}

func (c *Classrom) AddStudent(name string, age int) {
	newStuden := Studend8{
		Name: name,
		Age:  age,
	}
	c.classroom = append(c.classroom, newStuden)
}
func (c *Classrom) AverageAge() float64 {
	totalAge := 0
	for _, stutent := range c.classroom {
		totalAge += stutent.Age
	}
	return float64(totalAge) / float64(len(c.classroom))
}

// 9. Склад товаров
type Product9 struct {
	name     string
	quantity int
}
type Warehouse struct {
	products []Product9
}

func (w *Warehouse) AddProduct(name string, quantity int) {
	newWarehouse := Product9{
		name:     name,
		quantity: quantity,
	}
	w.products = append(w.products, newWarehouse)
}
func (w Warehouse) TotalQuantity() int {
	total := 0
	for _, p := range w.products {
		total += p.quantity
	}
	return total
}

// 10. Заметки и метки
type Note struct {
	content string
	tags    []string
}
type Notebook struct {
	notes []Note
}

func (n *Notebook) AddNote(content string, tags []string) {
	n.notes = append(n.notes, Note{
		content: content,
		tags:    tags,
	})
}
func (n Notebook) NotesWithTag(tag string) []Note {
	var result []Note
	for _, note := range n.notes {
		for _, t := range note.tags {
			if t == tag {
				result = append(result, note)
				break
			}
		}
	}
	return result
}

// 11. Управление зарплатами
type Employee struct {
	name   string
	salary float64
}
type Payroll struct {
	employees []Employee
}

func (p *Payroll) AddEmployee(name string, salary float64) {
	p.employees = append(p.employees, Employee{
		name:   name,
		salary: salary,
	})
}
func (p Payroll) AverageSalary() float64 {
	totalSalary := 0.0
	for _, emp := range p.employees {
		totalSalary += emp.salary
	}
	return totalSalary / float64(len(p.employees))
}

// 12. Разделение бюджета
type Budget struct {
	amount float64
}

func (b *Budget) Allocate(amount float64) bool {
	if amount <= b.amount {
		b.amount -= amount
		return true
	}
	return false
}
func (b *Budget) Remaining() float64 {
	return b.amount
}

// 13. Отслеживание заказов
type Order struct {
	id          int
	totalAmount float64
}
type OrderManager struct {
	orders []Order
}

func (o *OrderManager) AddOrder(id int, totalAmount float64) {
	o.orders = append(o.orders, Order{
		id:          id,
		totalAmount: totalAmount,
	})
}
func (o *OrderManager) TotalOrdersAmount() float64 {
	amount := 0.0
	for _, ord := range o.orders {
		amount += ord.totalAmount
	}
	return amount
}

// 14. Управление аккаунтами
type Account14 struct {
	balance float64
}

func (a *Account14) Deposit(amount float64) {
	a.balance += amount
}
func (a *Account14) Withdraw(amount float64) bool {
	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}
func (a *Account14) GetBalance() float64 {
	return a.balance
}

// 15. Операции над строками
type Text struct {
	content string
}

func (t *Text) AddText(text string) {
	t.content += text
}
func (t *Text) ReplaceWord(oldWord, newWord string) {
	t.content = strings.ReplaceAll(t.content, oldWord, newWord)
}
func (t *Text) Length() int {
	return len(t.content)
}

// 16. Управление списком покупок
type ShoppingItem struct {
	name  string
	price float64
}
type ShoppingList struct {
	shoppingItems []ShoppingItem
}

func (s *ShoppingList) AddItem(name string, price float64) {
	s.shoppingItems = append(s.shoppingItems, ShoppingItem{
		name:  name,
		price: price,
	})
}
func (s *ShoppingList) TotalPrice() float64 {
	total := 0.0
	for _, shop := range s.shoppingItems {
		total += shop.price
	}
	return total
}

// 17. Учет времени
type Event struct {
	name string
	date string
}
type Calendar struct {
	events []Event
}

func (e *Calendar) AddEvent(name string, date string) {
	e.events = append(e.events, Event{
		name: name,
		date: date,
	})
}
func (e *Calendar) EventsAfterDate(date string) []Event {
	var result []Event
	referenceData, _ := time.Parse("2006-01-02", date)
	for _, event := range e.events {
		eventDate, _ := time.Parse("2006-01-02", event.date)
		if eventDate.After(referenceData) {
			result = append(result, event)
		}
	}
	return result
}

// 18. Управление заказами в магазине
type Order18 struct {
	orderID  int
	product  string
	quantity int
}
type Store struct {
	order18s []Order18
}

func (s *Store) AddOrder(orderID int, product string, quantity int) {
	s.order18s = append(s.order18s, Order18{
		orderID:  orderID,
		product:  product,
		quantity: quantity,
	})
}
func (s *Store) TotalQuantityByProduct(product string) int {
	total := 0
	for _, order := range s.order18s {
		if order.product == product {
			total += order.quantity
		}
	}
	return total
}

// 19. Расчет стоимости поездки
type Trip struct {
	distance    float64
	costPerMile float64
}

func (t *Trip) SetCostPerMile(cost float64) {
	t.costPerMile = cost
}
func (t *Trip) TotalCost() float64 {
	return t.distance * t.costPerMile
}
