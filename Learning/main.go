package main

import "fmt"


var mathList map[string]func()

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int){
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial", i)

	zeroval(i)
	fmt.Println("zero Variable", i)
	fmt.Println("Zero Pointer", i)
	zeroptr(&i)
	fmt.Println("Zero Pointer", i)

	fmt.Println("Pointer", &i)

	mathList = map[string]func(){}
}
