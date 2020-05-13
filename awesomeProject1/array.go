package main

import "fmt"

func main(){
	var names [3]string

	names[0] = "eko"
	names[1] = "imam"
	names[2] = "joko"

	for i:=0;i<3;i++{
		fmt.Println(names[i])
	}

	fmt.Println(names[1])
	fmt.Println(names[0])
	fmt.Println(names[2])
}