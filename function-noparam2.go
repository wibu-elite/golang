package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func RandomNumber(){
	rand.Seed(time.Now().UTC().Unix())
	
	for i := 1; i < 5; i++{
		fmt.Println("Random Number ke", i,":", rand.Intn(7))
	}
}

func main() {
	fmt.Print("\x1bc")
	fmt.Println(`Function With Paramater
	`)
	var number int
	for true{
		fmt.Printf("Masukkan Angka: ")
		fmt.Scanln(&number)

		if number == 1{
		RandomNumber()
		}else if number == 2{
		fmt.Println("Anda Telah Keluar Program")
		os.Exit(2)
		}else {
		fmt.Println("Pilihanmu Tidak Ada")
		}
	}
	
}