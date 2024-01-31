package task26

/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
 Функция проверки должна быть регистронезависимой.

Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов в строке (регистронезависимо)
func areAllCharactersUnique(input string) bool {
	seen := make(map[rune]struct{})
	input = strings.ToLower(input) // Приводим к нижнему регистру для регистронезависимой проверки

	for _, char := range input {
		if _, found := seen[char]; found {
			return false // Найден повторяющийся символ
		}
		//struct использует для экономии памяти
		seen[char] = struct{}{}
	}
	return true // Все символы уникальны
}

func Run() {
	input1 := "abcd"
	fmt.Println(input1, "-", areAllCharactersUnique(input1)) // Output: abcd - true

	input2 := "abCdefAaf"
	fmt.Println(input2, "-", areAllCharactersUnique(input2)) // Output: abCdefAaf - false

	input3 := "aabcd"
	fmt.Println(input3, "-", areAllCharactersUnique(input3)) // Output: aabcd - false
}
