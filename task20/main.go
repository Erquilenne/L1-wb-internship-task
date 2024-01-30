package task20

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWordsInString(input string) string {
	words := strings.Fields(input)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func Run() {
	str := "snow dog sun"
	reversed := reverseWordsInString(str)
	fmt.Println("base - ", str, "\nreversed - ", reversed) // Output: sun dog snow
}
