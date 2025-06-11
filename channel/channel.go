package main

import (
	"fmt"
	"sync"
)

func main(){
	mt:=&sync.Mutex{}
	mych:=make(chan int,5)
    mt.Lock()
	mych<-10
   mt.Unlock()
	fmt.Println(<-mych)

}