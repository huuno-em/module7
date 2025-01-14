package calc

import "fmt"

//Вам предстоит доработать калькулятор, который вы делали в предыдущем модуле, опираясь на полученные знания о структурах и функциях. Создайте пакет calc, содержащий неэкспортируемую структуру calculator. Напишите экспортируемую функцию-конструктор для создания экземпляра структуры calculator. Добавьте экспортируемый (публичный) метод Calculate для структуры calculator. В качестве параметров метод должен принимать 2 числа типа float64 и строчный оператор (+, -, *, /). Возвращать метод должен значение типа float64. Добавьте неэкспортируемые (приватные) методы для каждой операции (сложения, вычитания, умножения, деления). Каждый приватный метод должен принимать на вход 2 числа типа float64 и возвращать значение типа float64. (В методе деления чисел должны быть проверка делителя на равенство 0). Метод Calculate должен в зависимости от переданного оператора вызывать один из приватных методов. (Если в качестве оператора передан +, то должен быть вызван приватный метод для сложения чисел). Cоздайте пакет main c функцией main(). Внутри функции main необходимо из консоли считать 2 числа и оператор (как вы это делали в задании предыдущего модуля). Затем создать экземпляр структуры calculator из пакета calc и вызвать метод Calculate, передав ему полученные из консоли значения. Полученный из метода Calculate результат нужно распечатать в консоль.

type calculator struct {
	num1, num2 float64
	operator   string
}

func NewCalc() calculator {
	return calculator{}
}

func (c calculator) Calculate(n1 float64, n2 float64, o string) float64 {
	c.num1 = n1
	c.num2 = n2
	c.operator = o
	switch c.operator {
	case "+":
		return c.addition()
	case "-":
		return c.subtraction()
	case "*":
		return c.multiplication()
	case "/":
		return c.division()
	default:
		fmt.Println("Error: Invalid operator")
		return 0
	}
}

func (c *calculator) multiplication() float64 {
	return c.num1 * c.num2
}

func (c *calculator) addition() float64 {
	return c.num1 + c.num2
}

func (c *calculator) subtraction() float64 {
	return c.num1 - c.num2
}

func (c *calculator) division() float64 {
	if c.num2 != 0 {
		return c.num1 / c.num2
	} else {
		fmt.Println("Error: Division by 0")
		return 0
	}
}
