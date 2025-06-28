package basics

import (
	"fmt"
	"math"
	"testing"
)

// why go math prefer float64 ?
type shape interface {
	name() string
	area() float64
	perim() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Phi * c.radius
}

func (c Circle) name() string {
	return "Circle"
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (r Rect) name() string {
	return "Rect"
}

func meausre(s shape) {
	fmt.Println(s.name())
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func TestIF(t *testing.T) {
	c := Circle{10}
	r := Rect{width: 10, height: 20}

	meausre(c)
	meausre(r)

}
