package main

import "fmt"

func main (){
	fmt.Printf("\x1bc") //Clear screen
	var name string = "Key"
	var age int = 17
	fmt.Println(name)
	fmt.Println(age)

	//Variabel without 'var'
	name1 := "Kun"
	age1 := 17
	fmt.Println(name1)
	fmt.Println(age1)

	//Multiple variabel
	var(
		name2 = "Tempest"
		age2 = 17
	)
	fmt.Println(name2)
	fmt.Println(age2)

}