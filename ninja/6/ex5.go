package ninja6

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

type Figure interface {
	area()
}

func (c Circle) area() {
	area := 2 * math.Pi * c.Radius
	fmt.Println("Área do Circle: ", area)
}

func (s Square) area() {
	area := s.Side * s.Side
	fmt.Println("Área do Square: ", area)
}

func Ex5() {
	circle := Circle{
		Radius: 2,
	}

	square := Square{
		Side: 5,
	}

	ShowInfo(circle)
	ShowInfo(square)
}

func ShowInfo(f Figure) {
	f.area()
}
