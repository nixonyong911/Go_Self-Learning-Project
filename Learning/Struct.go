package main

import "fmt"

type human struct{
	name string
	age int
}

func main() {
	fmt.Println(human{"bob", 20})
	fmt.Println(human{"Alice", 55})
	person1 := human{"flyer", 22}
	fmt.Println(person1.name)
	fmt.Println(person1)
	person1.name = "new name"
	fmt.Println(person1)
	person2 := &person1
	fmt.Println(person2.age)
	chgAge(&person2.age)
	fmt.Println(person2.age)
	chgAge(&person1.age)
	fmt.Println(person1.age)

}

func chgAge(age *int){
	*age = 50
}
