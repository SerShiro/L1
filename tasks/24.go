package tasks

import (
	"fmt"
	"math"
)

// делаем структуру с полями x y
type point struct {
	x, y float64
}

// конструктор
func newPoint(x float64, y float64) point {
	return point{x: x, y: y}
}

// вывод на экран координат
func (p *point) printCoordinate() {
	fmt.Printf("x = %.2f, y = %.2f\n", p.x, p.y)
}

// DistanceBetweenPoints
// метод для точки, который принимает аргументом еще одну точку, и по формуле рассчитываем расстояние между нами
func (p *point) DistanceBetweenPoints(p1 point) float64 {
	return math.Sqrt(math.Pow(p.x-p1.x, 2) + math.Pow(p.y-p1.y, 2))
}

func Task24() {
	point1 := newPoint(4, 5)
	point2 := newPoint(-1, 0)
	fmt.Printf("Координаты первой точки: ")
	point1.printCoordinate()
	fmt.Printf("Координаты второй точки: ")
	point2.printCoordinate()
	fmt.Printf("Расстояние между точками = %.2f\n", point1.DistanceBetweenPoints(point2))
}
