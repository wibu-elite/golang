package main

import "fmt"

func main() {
	fmt.Println("Tipe Data Declaration")

	type NoKTM int
	
	var myPass NoKTM
	fmt.Printf("Masukkan Nomor KTM :")
	fmt.Scanln(&myPass)

	fmt.Println("Your NPM : ", myPass)
}