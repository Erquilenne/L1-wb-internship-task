package task1

/*Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
в структуре Action от родительской структуры Human (аналог наследования).
*/

/* В Go вместо термина "наследование" используется термин "встраивание" (embedding) для создания подобного эффекта.
Встраивание позволяет структуре включать в себя другую структуру, тем самым получая доступ к её полям и методам.
*/

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

// Метод для структуры Human
func (h *Human) Introduce() {
	fmt.Printf("Hi, I'm %s, a %d-year-old %s.\n", h.Name, h.Age, h.Gender)
}

// Определение структуры Action с встраиванием структуры Human
type Action struct {
	Human // встраивание структуры Human
}

// Дополнительный метод для структуры Action
func (a *Action) DoSomething() {
	fmt.Printf("%s is doing something!\n", a.Name)
}

func Run() {
	// Создание объекта Action
	person := Action{
		Human: Human{
			Name:   "John",
			Age:    25,
			Gender: "Male",
		},
	}

	// Вызов метода Introduce из структуры Human
	person.Introduce()

	// Вызов метода DoSomething из структуры Action
	person.DoSomething()
}
