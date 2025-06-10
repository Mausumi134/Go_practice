package main

import "fmt"

func main(){
	var n int
    fmt.Print("enter number of entries: ")
	fmt.Scan(&n)
	arraymap:=make(map[string]int)
   
	for i:=0;i<n;i++{
		var key string
		var val int
		fmt.Scan(&key)
		fmt.Scan(&val)

		arraymap[key]=val
	}

	for k,v:=range(arraymap){
		fmt.Println("key is: ",k,"-value is: ",v)
	}

	slicemap:=[]map[string]string{
       {"name":"mausumi","city":"odisha"},
	}

	slicemap = append(slicemap,map[string]string{
		"name":"raju",
		"city":"kaju",
	} )

	for i,m:=range slicemap{
        fmt.Printf("Person %d:\n", i+1)
		for k, v := range m {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
	}


