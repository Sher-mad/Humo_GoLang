package main

import "fmt"

// Факториал
func main() {
	// var z int = 1256
	// b := &z
	// var c int = 1257
	// n := &c
	// var v int = 1258
	// m := &v

	// fmt.Println("b = ", b)
	// fmt.Println("n = ", n)
	// fmt.Println("m = ", m)
	// fmt.Println(" bb = ", *b)

	// var x int = 10
	// var ptr *int
	// ptr = &x
	// fmt.Println("Значение х = : ", x)
	// fmt.Println("Адрес = ", &x)
	// fmt.Println("Адрес ptr = ", ptr)
	// fmt.Println("Выводит значения ptr = ", *ptr)

	// x :=6
	// var v *int

	// Напишите функцию updateValue, которая принимает указатель на целое число и увеличивает его значение на 10.
	// n := 5
	// fmt.Println("before: ", n)
	// ubdateValue(&n)
	// fmt.Println("after: ", n)

	// Напишите функцию swap, которая меняет местами значения двух переменных с использованием указателей.
	// s := 3
	// d := 5
	// fmt.Println("before", s, d)
	// swap(&s, &d)
	// fmt.Println("after", s, d)

	var candidate1Votes, candidate2Votes int
	candidate1VotesPtr := &candidate1Votes
	candidate2VotesPtr := &candidate2Votes
	var (
		cand1 = "Кандидат 1"
		cand2 = "Кандидат 2"
	)
	vote(&cand1, candidate1VotesPtr)
	vote(&cand2, candidate2VotesPtr)
	vote(&cand1, candidate1VotesPtr)
	vote(&cand1, candidate1VotesPtr)
	fmt.Println("Победитель выборов:", winner(candidate1VotesPtr, candidate2VotesPtr))



}

// Напишите функцию updateValue, которая принимает указатель на целое число и увеличивает его значение на 10.
func ubdateValue(p *int) {
	*p = *p + 10
	fmt.Println(p)
}

// Напишите функцию swap, которая меняет местами значения двух переменных с использованием указателей.
func swap(x, y *int) {
	*x, *y = *y, *x
}

// func averageScore(scoresPtr *[]int)float64{
// 	return 0.0


// Middle Level
//3 Напишите программу для учета результатов тестирования студентов. Реализуйте функции для добавления результатов теста и вывода среднего балла.

func vote(candidatePtr *string, votesPtr *int) {
	*votesPtr++
}
func winner(candidate1VotesPtr *int, candidate2VotesPtr *int) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Кандидат 1"
	} else if *candidate1VotesPtr < *candidate2VotesPtr {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}

// Напишите программу для учета расходов на семейный бюджет. Реализуйте функции для добавления новой записи расходов и вывода общей суммы расходов.
func addExpense(totalPtr *float64, amount float64) {
	*totalPtr += amount
}
func pritTotalExpenses(totalPtr *float64) {
	fmt.Printf("Общая сумма расходов: %.2f\n", *totalPtr)

}

// Middle Level

// Напишите программу для учета оценок по предметам студентов. Реализуйте функции для добавления оценки по предмету и вывода среднего балла.
func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}
func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemistryPtr
	return float64(sum) / 3.0

}

// n!=(n*(n-1)*(n-2)...1)/(n-1)
func factorial(n int) int {
	result := 1
	for i := 1; i < n; i++ {
		result *= i
	}
	return result
}

func ubdate2(num int) {
	sum := num + 10
	fmt.Println("Значания sum", sum)
}
