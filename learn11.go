package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\x1bc")
	fmt.Println("Program Input Output Sederhana")
	
	var(
		name string
		npm int
		status string
	)
	fmt.Printf("Add Your Name : ")
	fmt.Scanln(&name)

	fmt.Printf("Add Your NPM : ")
	fmt.Scanln(&npm)
	
	fmt.Printf("Single or Married (Yes/No) ")
	fmt.Scanln(&status)


	fmt.Println("Your Name :", name)
	fmt.Println("Your NPM :", npm)

	status = strings.ToLower(status)

	if status == "no"{
		fmt.Println("You are Single")
	} else {
		fmt.Println("You Have Been Married")
	}
	
}
