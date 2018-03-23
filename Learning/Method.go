package main

import "fmt"


type rect struct{
	width, height int
}

func main() {
	woho := rect{50, 20}

	fmt.Println(woho.area())
	fmt.Println(woho.para())
}

func (r rect) area() int{
	return r.height * r.width
}

func (r rect) para() int{
	return 2 * r.height + 2 * r.width
}