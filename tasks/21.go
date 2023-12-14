package tasks

import "fmt"

// Rectangle
// у нас есть интерфейс который должен имплементиться нашей структурой
type Rectangle interface {
	getArea() int
}

// Rect
// есть стурктура, которая реализует данный интерфейс
type Rect struct {
	width, height int
}

func (r *Rect) getArea() int {
	fmt.Printf("Я метод прямоугольника и я вывожу width * height: %d * %d = %d\n", r.width, r.height, r.width*r.height)
	return r.width * r.height
}

func NewRect(width int, height int) *Rect {
	return &Rect{width: width, height: height}
}

// Square
// есть структура, которую мы не можем изменить, но нужно чтобы она тоже реализовывала интерфейс
type Square struct {
	side int
}

func (s *Square) Area() int {
	fmt.Printf("Я метод квадрата и я вывожу side * side: %d * %d = %d\n", s.side, s.side, s.side*s.side)
	return s.side * s.side
}

// SquareAdapter
// создаем адаптер для нашей структуры
type SquareAdapter struct {
	*Square
}

func NewSquareAdapter(square *Square) *SquareAdapter {
	return &SquareAdapter{Square: square}
}

// реализуем интерфейс
func (sa *SquareAdapter) getArea() int {
	return sa.Square.Area()
}

func Task21() {
	rect := NewRect(4, 5)
	rect.getArea()
	adapter := NewSquareAdapter(&Square{3})
	adapter.getArea()
}
