package main

import "fmt"

func main(){
	fmt.Printf("\x1bc")
	var value32 int32 = 1000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)

	//konversi beda tipe data
	me := "Key"
	var e byte = me[0]
	var eString = string(e)
	fmt.Println(e)
	fmt.Println(eString)


}