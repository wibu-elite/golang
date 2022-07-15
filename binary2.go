package main

import (
	"fmt"
)

var (
	banyakData int
	urutkan    int
	cariData   int
	Hasil      int
)


// Fungsi Mengurutkan Bilangan
func BinarySearch(banyakData int, urutkan int, cariData int) {
	fmt.Printf("Masukkan Bnayak Data: ")
	fmt.Scanln(&banyakData)

	var urut int
	var value = make([]int, 254)
	for i := 0; i < banyakData; i++ {
		fmt.Print("Masukkan Data Ke", i+1,": ")
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
	fmt.Printf(`
Masukkan Data yang Ingin Dicari: `)
	fmt.Scanln(&cariData)
	fmt.Println(cariData)

	for i := 0; i < len(value); i++{
		low := 0
		high := len(value)
		mid := (low + high) / 2
		if cariData < value[mid] {
			high = mid - 1
		} else if cariData > value[mid] {
			high = mid + 1
		} else if cariData == value[mid] {
			
		}
	}
}
func main() {
	var banyakData int
	fmt.Printf("\x1bc")
	fmt.Println(`Program Binary Searching Dengan Inputan Dari User
	`)

	BinarySearch(banyakData, urutkan, cariData)
}