package task21

//Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// Интерфейс цели
type RussianSpeaker interface {
	SayHelloInRussian() string
}

// Адаптируемая структура
type AmericanSpeaker struct{}

func (a *AmericanSpeaker) GreetInEnglish() string {
	return "Hello"
}

// Адаптер
type RussianSpeakerAdapter struct {
	americanSpeaker *AmericanSpeaker
}

func (r *RussianSpeakerAdapter) SayHelloInRussian() string {
	englishGreeting := r.americanSpeaker.GreetInEnglish()
	// Простой пример адаптации - замена "Hello" на "Привет"
	russianGreeting := "Привет"
	return fmt.Sprintf("%s, %s", russianGreeting, englishGreeting)
}

func Run() {
	american := &AmericanSpeaker{}
	adapter := &RussianSpeakerAdapter{americanSpeaker: american}
	fmt.Println(adapter.SayHelloInRussian()) // Output: Привет, Hello
}
