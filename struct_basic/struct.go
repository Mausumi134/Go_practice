package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := new(person)
	p.Name = "mausumi"
	p.Age = 20

	fmt.Println(p.Name, p.Age)

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error")
		return
	} else {
		fmt.Println(string(jsonData))  
	}
}
