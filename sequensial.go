package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("\x1bc")
	fmt.Println(`PROGRAM SEDERHANA SEARCHING SEQUENSIAL MIFTA N R 10
	`)
	text := `
List Data :                                           
1.  Januari                                          
2.  Februari
3.  Maret
4.  April
5.  Mei
6.  Juni
7.  Juli
8.  Agustus
9.  September
10. Oktober
11. November
12. Desember
	`
	fmt.Println(text)

	var (
		month = [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli",
			"Agustus", "September", "Oktober", "November", "Desember"}

		search string
	)
	fmt.Printf("Masukka Data yang Ingin Dicari: ")
	fmt.Scanln(&search)

	// }
	for i := 0; i < len(month); i++ {
		if strings.EqualFold(month[i], search) {
			fmt.Println("Data Ditemukan Pada Index Ke: ", i)
			break
		}
		if i == (len(month) - 1) {
			fmt.Println("Tidak Ada Data Ditemukan")
			// break
		}
	}
	// for i := 0; i < 12; i++ {
	// 	// search := strings.ToLower(search)
	// 	if month[i] != search {
	// 		fmt.Println("Tidak Ada Data Ditemukan")
	// 		return
	// 	}
	// }
}
