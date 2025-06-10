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
	


	// check if item exist or not

   found:=false
	for _,num:=range(arr){
		if num==4{
			found=true
			break
		}
	}
if found {
    fmt.Println("num exist")
} else {
    fmt.Println("not exist")
}


// find and delete an item in an array

 result:=deleteArr(arr,3)
 fmt.Println(result)

//  append into the array
add(&arr)
fmt.Print(arr)
}

