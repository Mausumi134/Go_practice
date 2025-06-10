package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p:= new(person)

	p.name="mausumi"
	p.age=20

	fmt.Println(p.name,p.age)

	
}