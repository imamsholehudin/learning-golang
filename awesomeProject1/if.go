package main

import "fmt"

func main(){

	for i:=1; i<=30;i++ {
		if i%2 ==1{
			fmt.Println(i,"Ganjil")

		}	else if i%2 == 0 {
			fmt.Println(i,"Genap")
		}
	}

	for j:=1; j<30;j++{
		if j% 3 == 0 && j%5 ==0{
			fmt.Println("FooBar")
		}else if j % 3 == 0{
			fmt.Println("Foo")

		}else if j % 5 == 0 {
			fmt.Println("Bar")
		}else {
			fmt.Println(j)

		}	}
}
