package main

import (
	"fmt"
)

// var ()

// Fungsi Mengurutkan Bilangan
func main() {
	fmt.Printf("\x1bc")
	var (
		urut       int
		value      = make([]int, 254)
		banyakData int
		mid        int
		cariData   int
		// i = 0
		// Hasil      int
	)
	fmt.Printf("Masukkan Bnayak Data: ")
	fmt.Scanln(&banyakData)
	for i := 0; i < banyakData; i++ {
		fmt.Print("Masukkan Data Ke", i+1, ": ")
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
	low := 0
	high := banyakData
	// fmt.Println(high)
	for low <= high {
		mid = (high + low) / 2
		if cariData < value[mid] {
			high = mid - 1
		}
		if cariData > value[mid] {
			low = mid + 1
		}
		if value[mid] == cariData {
			fmt.Print("Data Ditemukan Pada Index Ke: ", mid-1)
			return
		}
	}
	if value[mid] != cariData {
		fmt.Println("Data Tidak Ditemukan")
		return
	}

	// for i := 0; i < banyakData; i++{
	// 	if value[mid] == cariData{
	// 		fmt.Print("Data", cariData, "Ditemukan Pada ", mid)
	// 	}
	// }
}
 