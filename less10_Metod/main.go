package main

import "fmt"

func main() {
	rect := Rectangle{Width: 10, heigth: 5}
	area := rect.Acer()
	fmt.Printf("Area : %.3f", area)

}
