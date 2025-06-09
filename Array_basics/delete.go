package main


func deleteArr(arr[] int,num int )[]int{

	result:=[]int{}

	for _,v:=range(arr){
		if(v!=num){
			result = append(result, v)
		}
	}
	return result
}