package main

import "fmt"

func main() {
	fmt.Println("\x1bc")
	fmt.Println(`PROGRAM SEARCHING BINARY MIFTA 10
	`)
	var (
		capaC      = [7]int{}
		valueL int = 0
		valueR int = 7
		search int
	)
	for i := 0; i < (len(capaC)); i++ {
		fmt.Printf("Masukkan Data: ")
		fmt.Scanln(&capaC[i])
	}
	fmt.Printf("Masukkan Data yang Ingin Dicari: ")
	fmt.Scanln(&search)

	for i := 0; i < (len(capaC)); i++ {
		valueM := (valueL + valueR) / 2
		fmt.Println(valueM)
		if capaC[valueM] == search {
			fmt.Println("Data Ditemukan Pada Index ke :", i)
			break
		}

	}

}
