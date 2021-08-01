package main

import (
	"fmt"
	"log"
)

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}
func main() {
	divideByZero()
	x := int(1)
	y := int(2)
	fmt.Println("Before swap x", x)
	fmt.Println("Before swap y", y)
	swap(&x, &y)
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println("deref", *&x)

	fmt.Println("deref", *&y)
	fmt.Println("After swap x", x)
	fmt.Println("After swap y", y)

}

//fmt.Println(10 / 0)
func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println("output", divide(3, 0))
}
func divide(num1, num2 int) int {
	return num1 / num2

}
