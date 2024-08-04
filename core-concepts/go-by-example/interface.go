package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectArea struct {
	width, height float64
}
type circleArea struct {
	radius float64
}

func (r rectArea) area() float64 {
	return r.width * r.height
}
func (r rectArea) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circleArea) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circleArea) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectArea{width: 3, height: 4}
	c := circleArea{radius: 5}

	measure(r)
	measure(c)
}
