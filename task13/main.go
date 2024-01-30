package task13

//Поменять местами два числа без создания временной переменной.

import "fmt"

func Run() {
	a, b := 5, 10

	fmt.Println("До обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("После обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
