package main

import "fmt"

// slice of struct in golang

type Person struct{
	name string
	age int
}
func main(){
   boy:=[]Person{
	{"mausumi",20},
	{"john",19},
   }
   boy=append(boy, Person{"roy",15})
   for _,p:=range boy{
	fmt.Println("name is: ",p.name,",age is: ",p.age)

   }

   index:=1

   if index>=0 && index<len(boy){
	boy = append(boy[:index],boy[index+1:]...)
   }
   for _,p:=range boy{
	fmt.Println("name is: ",p.name,",age is: ",p.age)

   }
}