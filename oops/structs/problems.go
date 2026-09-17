package structs

import (
	"fmt"
	"math"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.items) - 1
	v := s.items[index]
	s.items = s.items[:index]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(p2 Point) float64 {
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func RunStructs() {
	fmt.Println("Struct rollin")

	var strStack Stack[string]

	strStack.Push("hello")
	strStack.Push("hi")
	strStack.Push("Aashish")

	fmt.Println(strStack)
	fmt.Println(strStack.Pop())
	fmt.Println(strStack.Peek())
	fmt.Println(strStack)

	pointA := Point{X: 5, Y: 10}
	pointB := Point{X: 10, Y: 20}

	fmt.Println(pointA.Distance(pointB))

}
