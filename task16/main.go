package task16

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
	"sort"
)

func quickSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func Run() {
	arr := []int{3, 2, 5, 1, 4}
	quickSort(arr)
	fmt.Println(arr)
}
