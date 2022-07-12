package main

import "fmt"

func main() {
	fmt.Println("\x1bc")
	fmt.Println("Program Operasi Matematika")

	var(
		a int
		b int
		result int
	)
	fmt.Printf("Masukkan Nilai A: ")
	fmt.Scanln(&a)
	fmt.Printf("Masukkan Nilai B: ")
	fmt.Scanln(&b)
	result = a*b

	fmt.Println(result)
}