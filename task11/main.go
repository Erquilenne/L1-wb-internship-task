package task11

//Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

func intersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	var intersection []int

	for _, num := range set1 {
		setMap[num] = true
	}

	for _, num := range set2 {
		if setMap[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}

func Run() {
	set1 := []int{7, 2, 6, 12, 5}
	set2 := []int{13, 4, 12, 6, 7}
	result := intersection(set1, set2)
	fmt.Println(result)
}
