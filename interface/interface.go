package main

import "fmt"

type Ancestor interface {
	Add(A int, B int) int
	Substract(A int, B int) int
}

type Person struct {
	A int
	B int
}

func (p Person) Add(A int, B int) int {
	return A + B
}
func (p Person) Substract(A int, B int) int {
	return A - B
}
func main() {
	p := new(Person)
	p.A = 4
	p.B = 2

	c := p.Add(p.A, p.B)
	d := p.Substract(p.A, p.B)

	fmt.Println(c)
	fmt.Println(d)
}