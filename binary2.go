package main

import (
	"fmt"
)

var (
	banyakData int
	urutkan    int
	cariData   int
	Hasil int
)

// Fungsi Mengurutkan Bilangan
func UrutkanData(banyakData int, urutkan int) {
	fmt.Printf("Masukkan Bnayak Data: ")
	fmt.Scanln(&banyakData)

	var urut int
	var value = make([]int, 254)
	for i := 0; i < banyakData; i++ {
		fmt.Printf("Masukkan Data: ")
		fmt.Scanln(&value[i])
	}

	for x := 0; x < banyakData; x++ {
		for z := 0; z < banyakData; z++ {
			if value[z] > value[z+1] {
				urut = value[z]
				value[z] = value[z+1]
				value[z+1] = urut
			}
		}
	}
	fmt.Println(`
Hasil Data Setelah Diurutkan: `)
	for i := 0; i < banyakData; i++ {
		fmt.Print(value[i+1], " ")
	}

}

//Fungsi

func BinarySeacrh(banyakData int, cariData int) {
	fmt.Println(banyakData)
}

func main() {
	var banyakData int
	fmt.Printf("\x1bc")
	fmt.Println(`Program Binary Searching Dengan Inputan Dari User
	`)
	// for 
	UrutkanData(banyakData, urutkan)


}

