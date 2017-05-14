package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Nob", 5})
	fmt.Println(person{name: "Nob", age: 5})
	fmt.Println(person{name: "Lan"})

	s := person{name: "love", age: 10000}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)
}
