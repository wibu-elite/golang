package main

import "fmt"

var(
	bil1 float64
	bil2 float64
	tipe string
)
func main() {
	fmt.Println("\x1bc")
	fmt.Println("My Function Multiplication")
	fmt.Printf("Masukkan Bilangan 1 :")
	fmt.Scanln(&bil1)
	fmt.Printf("Masukkan Bilangan 2 :")
	fmt.Scanln(&bil2)
	fmt.Printf("Masukkan Jenis Operasi :")
	fmt.Scanln(&tipe)
	fCalc(bil1,bil2, tipe)

}
func fCalc(bil1 float64, bil2 float64, tipe string){
	if tipe == "-" {
		result := bil1 - bil2
		fmt.Println(result)
	}else if tipe == "+" {
		result := bil1 + bil2
		fmt.Println(result)
	}else if tipe == "/" {
		result := bil1 / bil2
		fmt.Println(result)
	}else if tipe == "*" {
		result := bil1 * bil2
		fmt.Println(result)
	} else {
		fmt.Println("Your Operation Wrong")
	}
}
