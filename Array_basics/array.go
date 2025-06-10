package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)
func reverse(arr []int)[]int{
	n:=len(arr)

	for i:=0;i<n/2;i++{
      arr[i],arr[n-1-i]=arr[n-1-i],arr[i]
	}
	return arr
}
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
fmt.Println(arr)

// convert the array into json
jsonarr:=[]string{"hello","world"}
jsonData,err:=json.Marshal(jsonarr)

if err!=nil{
	fmt.Println("error in converting into json")
	return
}else{
	fmt.Println(string(jsonData))
}
// sort array
sort.Ints(arr)
fmt.Println(arr)

// reverse array
arr=reverse(arr)
fmt.Println(arr)

// append one slice to another slice

  stringArr := []string{"Go"}
  stringArr = append(stringArr, jsonarr...)
fmt.Println(stringArr)

// convert arr of int to string in go

strArr:=make([]string,len(arr))
for i,v:=range arr{
	strArr[i]=strconv.Itoa(v)
}
fmt.Println(strArr)
resultarr:=strings.Join(strArr,",")
fmt.Println(resultarr)
}

