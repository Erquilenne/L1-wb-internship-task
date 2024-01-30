package task14

import (
	"fmt"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Это int:", v)
	case string:
		fmt.Println("Это string:", v)
	case bool:
		fmt.Println("Это bool:", v)
	case chan int:
		fmt.Println("Это канал для передачи int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func Run() {
	var v1 interface{} = 42
	getType(v1)

	var v2 interface{} = "hello"
	getType(v2)

	var v3 interface{} = true
	getType(v3)

	ch := make(chan int)
	var v4 interface{} = ch
	getType(v4)
}
