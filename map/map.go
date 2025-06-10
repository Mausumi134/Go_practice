package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stringmap := map[string]int{
		"apple":  10,
		"mango":  5,
		"banana": 4,
	}

	_, exist := stringmap["apple"]

	if exist {
		fmt.Println("existed")
	}else{
		fmt.Println("not exist")
	}

	stringmap["apple"]=20

	for k,v := range stringmap{
		fmt.Println("key=>value",k,v)
	}

	// convert map to json
	jsonData,err:=json.Marshal(stringmap)
	if err!=nil{
		fmt.Println("error in converting to json")

	}else{
		fmt.Println(string(jsonData))
	}

	// convert json to map
   
	var result map[string]int
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("error in unmarshalling json")
	} else {
		for k, v := range result {
			fmt.Println("key=>value", k, v)
		}
	}

	// delete or remove key from map

	delete(stringmap,"mango")
    for k,v := range stringmap{
		fmt.Println("key=>value",k,v)
	}
}