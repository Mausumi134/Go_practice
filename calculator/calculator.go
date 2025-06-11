package main

import "fmt"

func main(){
	var a,b float64
	var op string
	fmt.Println("Enter first number: ")
    
	fmt.Scan(&a)

	fmt.Println("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Enter the operation: ")

	fmt.Scan(&op)

	switch op{
	case "+":
		fmt.Printf("%.2f+%.2f=%.2f\n",a,b,a+b)
	case "-":
		fmt.Printf("%.2f-%.2f=%.2f\n",a,b,a-b)
	case "*":
		fmt.Printf("%.2f*%.2f=%.2f\n",a,b,a*b)
	case "/":
		if b!=0{
		fmt.Printf("%.2f/%.2f=%.2f\n",a,b,a/b)
		}else{
			fmt.Println("Error division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}