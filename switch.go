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
	name = strings.ToLower(name)

	switch name {
	case "Key":
		fmt.Println("This is Key")
	case "Kun":
		fmt.Println("This is Kun")
	case "Tempest":
		fmt.Println("This is Tempest")
	default:
		fmt.Println("No Name")
	}
}
