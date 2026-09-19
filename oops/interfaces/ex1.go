package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t *Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	shapes := []Shape{
		&Circle{radius: 5},
		&Rectangle{width: 4, height: 5},
		&Triangle{a: 3, b: 4, c: 5},
	}

	fmt.Printf("Total Area: %.2f\n", TotalArea(shapes))
}
