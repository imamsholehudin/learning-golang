package main

import "fmt"

func main(){
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "imam"
	mahasiswa["1002"] = "sholeh"
	mahasiswa["1003"] = "udin"
	fmt.Println(mahasiswa)
	for nim, name := range mahasiswa{
		println("nim : "+nim+", Nama: "+name)
	}

	//map di dalam map

}