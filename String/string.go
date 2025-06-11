package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

// check if string is a number
  n:= "h e l l o"

  _,err:=strconv.ParseInt(n,10, 32)

  if err==nil{
	fmt.Println("it is a number")
  }else{
	fmt.Println("it is not a number")
  }

//   remove all whitespace from a string

var result string
for _,c:=range n{
	if(c!=' '){
       result+=string(c)
	}
}
fmt.Println(result)

n=strings.ReplaceAll(n," ","")

fmt.Println(n)

// Check if a string contains another string in GO


text:="Golang is fun"
sub:="fun"

if strings.Contains(text,sub){
	fmt.Println("Substring found")
}else{
	fmt.Println("Substring not found")
}

// Get all the words from a sentence

words:="hello world"

word:=strings.Split(words," ")

for _,w:=range word{
	fmt.Println(w)
}

// Check if string begins with a prefix in Go

if(strings.HasPrefix(words,"he")){
	fmt.Println("prefix found")
}else{
	fmt.Println("prefix not found")
}
// Check if string ends with a suffix in Go

if(strings.HasSuffix(words,"ld")){
	fmt.Println("suffix found")
}else{
	fmt.Println("suffix not found")
}
// Convert string to lowercase in Go

words=strings.ToLower(words)

fmt.Println(words)

words=strings.ToUpper(words)
fmt.Println(words)

// swap two strings 

string1:="hi"
string2:="bye"

string1,string2=string2,string1

fmt.Println(string1,string2)

fmt.Println("\033[1mThis is bold text\033[0m")
fmt.Println("\033[3mThis is italic text\033[0m")
fmt.Println("\033[35mThis is blue text\033[0m")
}