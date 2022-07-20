package main

import (
	"fmt"
	"strings"
)

func main() {
	print("\x1bc")
	var name string
	fmt.Printf("Masukkan Nama: ")
	fmt.Scanln(&name)
	// name = strings.ToLower(name)
	for true{
		switch{
		case strings.EqualFold(name, "key"):
			fmt.Println("This is Key")
			
		case strings.EqualFold(name, "kun"):
			fmt.Println("This is Kun")
			
		case strings.EqualFold(name, "tempest"):
			fmt.Println("This is Tempest")
			
		case strings.EqualFold(name, "keluar") :
			break
		default:
			fmt.Println("No Name")
		}
	}
	
}
