package main

import "fmt"
import "math"

func main() {
	r := rects{5,20}
	c := circle{5}

	measure(r)
	measure(c)
}

type geometry interface {
	area () float64
	perim () float64
}

type rects struct {
	height, width float64
}


type circle struct {
	radius float64
}

func (r rects) area() float64{
	return r.width * r.height
}

func (r rects) perim() float64{
	return 2 * r.width + 2 * r.height
}

func (r circle) area() float64{
	return math.Pi * r.radius * r.radius
}

func (r circle) perim() float64 {
	return 2 * math.Pi * r.radius
}

func measure (g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}