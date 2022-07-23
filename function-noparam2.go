package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	fmt.Println(`Random Data
	`)
	var(
		number string
		// code string
	) 
	for true{
		fmt.Printf("Masukkan Kode: ")
		fmt.Scanln(&number)
		number = strings.ToLower(number)

		if number == "1"{
		RandomNumber()
		}else if number == "2"{
		fmt.Println("Anda Telah Keluar Program")
		os.Exit(2)
		}else if number == "clear"{
			fmt.Print("\x1bc")
		}else {
		fmt.Println("Pilihanmu Tidak Ada")
		}
	}
}