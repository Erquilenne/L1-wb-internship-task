package task10

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

import "fmt"

func Run() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemperatures := make(map[int][]float64)

	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		groupedTemperatures[groupKey] = append(groupedTemperatures[groupKey], temp)
	}

	fmt.Println(groupedTemperatures)
}
