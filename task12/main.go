package task12

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import "fmt"

func Run() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, str := range sequence {
		set[str] = true
	}

	fmt.Println(set)
}
