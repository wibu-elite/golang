package main

import "fmt"

func main (){
	fmt.Printf("\x1bc")
	//variabel constant saat deklarasi harus ada valuenya
	const myname = "KunTempest"
	const myage = 17
	fmt.Println(myname)
	fmt.Println(myage)

	//Multiple Constant
	const(
		myname1 = "Tempest"
		myage1 = 17
	)
	fmt.Println(myname1)
	fmt.Println(myage1)
}