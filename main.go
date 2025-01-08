package main

//Напишите программу, которая считывает первое число, затем арифметический оператор (+, -, *, /), затем второе число, после чего, в зависимости от арифметического оператора, производит нужное действие и выводит строку с результатом.

import (
	"fmt"
)

func main() {

	var a, b, result float64
	var o string

	fmt.Scanln(&a)
	fmt.Scanln(&o)
	fmt.Scanln(&b)

	switch o {
	case "+":
		result = a + b
		fmt.Println(result)
	case "-":
		result = a - b
		fmt.Println(result)
	case "*":
		result = a * b
		fmt.Println(result)
	case "/":
		if b != 0 {
			result = a / b
			fmt.Println(result)
		} else {
			fmt.Println("Ошибка: деление на 0")
		}
	default:
		fmt.Println("Ошибка: неверная операция")
	}
}
