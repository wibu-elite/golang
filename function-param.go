package main

import (
	"fmt"
	"math/rand"
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
	RandomNumber()
}