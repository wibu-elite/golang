package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("\x1bc")
	fmt.Println("Learn Function of Strings")

	// 1. String contains() --> untuk mendeteksi string param kedua apakah bagian string param pertama
	//nilai kembaliannya boolean
	fmt.Println("STRING CONTAINS")
	var part1 string
	fmt.Printf("Input Yours :")
	fmt.Scanln(&part1)
	Stcontain := strings.Contains("Kun Tempest", part1)
	// Stcontain = strings.ToLower(Stcontain)
	fmt.Println(Stcontain)

	//2.strings.HasSuffix() -->Digunakan untuk deteksi apakah sebuah string (parameter pertama)
	// diakhiri string tertentu (parameter kedua).
	var last string
	fmt.Printf("Input last alfabet :")
	fmt.Scanln(&last)
	fmt.Println("STRING CONTAINS")
	Stsuffix := strings.HasSuffix("Kun Tempest", last)
	fmt.Println(Stsuffix)


}
