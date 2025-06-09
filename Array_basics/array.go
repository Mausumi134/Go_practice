package main

import "fmt"

func main(){
	// copy an array or slice in golang 
	var n int
    fmt.Print("Take the size of the array input: ")
	fmt.Scan(&n)

	arr:=make([]int,n)

	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}

    copyArr:=make([]int,n)

	for i:=0;i<n;i++{
		copyArr[i]=arr[i]
	}


	fmt.Println(copyArr)
	

}