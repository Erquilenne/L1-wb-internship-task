package task24

/*
Разработать программу нахождения расстояния между двумя точками,
 которые представлены в виде структуры Point с инкапсулированными
  параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

// Структура Point для представления точки
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func Run() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := point1.DistanceTo(point2)
	fmt.Println("Расстояние между точками:", distance)
}
