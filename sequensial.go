package main

import (
	"fmt"
	// "strings"
)

func main() {
	fmt.Printf("\x1bc")
	fmt.Println("PROGRAM SEDERHANA SEARCHING SEQUENSIAL MIFTA N R 10", "\n")
	fmt.Println(`List Data :`,"\n",
	`1.  Januari`,"\n",
	`2.  Februari`,"\n",
	`3.  Maret`,"\n",
	`4.  April`,"\n",
	`5.  Mei`,"\n",
	`6.  Juni`,"\n",
	`7.  Juli`,"\n",
	`8.  Agustus`,"\n",
	`9.  September`,"\n",
	`10. Oktober`,"\n",
	`11. November`,"\n",
	`12. Desember`,"\n")	

	var (
		month = [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli",
			"Agustus", "September", "Oktober", "November", "Desember"}

		search string
	)
	fmt.Printf("Masukka Data yang Ingin Dicari: ")
	fmt.Scanln(&search)

	// search := strings.ToLower(search)
	while(true){
		
	}
	for i := 0; i < 12; i++ {
		// search := strings.ToLower(search)
		if month[i] == search {
			fmt.Println("Data Ditemukan Pada Index Ke: ", i)
			return
		}
	}
	for i := 0; i < 12; i++ {
		// search := strings.ToLower(search)
		if month[i] != search {
			fmt.Println("Tidak Ada Data Ditemukan")
			return
		}
	}
}
