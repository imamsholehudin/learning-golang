package main

import "fmt"

func main(){
	i:=7
	switch i{
		case 0 : fmt.Println("nol")
		case 1 : fmt.Println("satu")
		case 2 : fmt.Println("dua")
		case 3 : fmt.Println("tiga")
		case 4 : fmt.Println("empat")
		case 5 : fmt.Println("lima")
		case 6 : fmt.Println("enam")
		case 7 : fmt.Println("tujuh")
		default : fmt.Println("default")
	}
}
