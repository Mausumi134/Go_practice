package main

import "fmt"

type Address struct{
	City string
	Zipcode int
}

type Person struct{
	Name string
	Age int
	Address Address
}
func main(){
p:=Person{
	Name:"mausmi",
	Age:20,
	Address: Address{
		City:"india",
		Zipcode: 11111,
	},
}
    fmt.Println("Name:", p.Name)
	 fmt.Println("Name:", p.Age)
    fmt.Println("City:", p.Address.City)
    fmt.Println("Zip Code:", p.Address.Zipcode)
}