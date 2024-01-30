package task17

import (
	"fmt"
	"sort"
)

//Реализовать бинарный поиск встроенными методами языка.

func Run() {
	arr := []int{3, 2, 5, 1, 4}
	target := 5
	result := binarySearch(arr, target)
	fmt.Println(result)
}

func binarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}
