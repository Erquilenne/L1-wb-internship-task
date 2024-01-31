package task22

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math"
	"math/big"
)

func Run() {
	calculate()
	calculateWithBig()
}

func calculate() {
	a := math.Pow(2, 21) // Значение переменной a
	b := math.Pow(2, 20) // Значение переменной b

	// Умножение
	multiplication := a * b
	fmt.Println("Умножение:", multiplication)

	// Деление
	division := a / b
	fmt.Println("Деление:", division)

	// Сложение
	addition := a + b
	fmt.Println("Сложение:", addition)

	// Вычитание
	subtraction := a - b
	fmt.Println("Вычитание:", subtraction)
}

func calculateWithBig() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil) // Значение переменной a
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil) // Значение переменной b

	// Умножение
	multiplication := new(big.Int).Mul(a, b)
	fmt.Println("Умножение with big:", multiplication)

	// Деление
	division := new(big.Rat).SetFrac(a, b)
	fmt.Println("Деление with big:", division)

	// Сложение
	addition := new(big.Int).Add(a, b)
	fmt.Println("Сложение with big:", addition)

	// Вычитание
	subtraction := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание with big:", subtraction)
}
