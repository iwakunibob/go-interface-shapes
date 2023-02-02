package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}
type square struct {
	side float64
}
type triangle struct {
	base   float64
	height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (sq square) area() float64 {
	return sq.side * sq.side
}
func (t triangle) area() float64 {
	return t.base * t.height * 0.5
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}
func (sq square) perim() float64 {
	return sq.side * 4
}
func (t triangle) perim() float64 {
	return t.base + t.height + (math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2)))
}

type shape interface {
	area() float64
	perim() float64
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{10}
	r1 := rectangle{4, 8}
	t1 := triangle{4, 8}
	s1 := square{4}
	shapes := []shape{&c1, &r1, t1, s1}
	fmt.Println("Shape          Area    Perimeter")
	for _, shape := range shapes {
		strStruct := fmt.Sprintf("%T", shape)
		fmt.Printf("%10s %10.2f %10.2f\n",
			strStruct[5:], shape.area(), shape.perim())
	}
	r2 := rectangle{5, 10}
	fmt.Println("Rectangle Area =", getArea(&r2))
	// Can also make more flexible using & for values and pointers
	fmt.Println("Triangle Area =", getArea(&t1))
}
