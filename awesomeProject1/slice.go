package main

import "fmt"

func main(){
	names:= [5]string{
		"Eko",
		"Budi",
		"Joko",
		"rudi",
		"Brian",
	}
	fmt.Println(names)

	slice1 := names[1:4]
	fmt.Println(slice1)

	sliceX:= []string{
		"Imam",
		"Sholeh",
		"Udin",
		"Kunco",
	}
	fmt.Println(sliceX)
	slice2:= append(slice1,"soko","cimong")
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string,3)
	copy(slice3, slice1)
	fmt.Println(slice3)
}
