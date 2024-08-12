package main

type Rectangle struct {
	Width  float64
	heigth float64
}

/*
 1. Книга и Автор
    •	Описание: Реализуйте структуру Book с полями title и author,
    а также метод GetDetails, который возвращает строку с деталями книги.
    Реализуйте структуру Library с массивом книг и метод AddBook,
    добавляющий книгу в библиотеку.
    •	Методы:
    •	GetDetails() string для структуры Book
    •	AddBook(book Book) для структуры Library
*/
func (r Rectangle) Acer() float64 {
	return r.Width * r.heigth
}
