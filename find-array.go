package main

import "fmt"

func main() {
	fmt.Println("\x1bc")
	fmt.Println("Find Word in Array")

	var isArray = make([]string, 254)
	var size int
	fmt.Printf("Input Size :")
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		fmt.Printf("Input Name : ")
		fmt.Scanln(&isArray[i])
	}

	for i := 0; i < size; i++ {
		fmt.Println("Value Index ke", (i), ":", isArray[i])
	}
}
