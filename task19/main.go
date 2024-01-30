package task19

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

import (
	"fmt"
)

// Функция для переворачивания строки с учётом символов Unicode
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Run() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Println("base - ", str, "\nreversed - ", reversed) // Output:абырвалг
}
