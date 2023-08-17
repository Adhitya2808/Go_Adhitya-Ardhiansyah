package main

import "fmt"

func main() {
	var N int16

	fmt.Print("Masukkan Nilai = ")
	fmt.Scan(&N)
	if N > 100 {
		fmt.Println("Nilai Invalid")
	} else if N >= 80 {
		fmt.Println("Grade A")
	} else if N >= 65 {
		fmt.Println("Grade B")
	} else if N >= 50 {
		fmt.Println("Grade C")
	} else if N >= 35 {
		fmt.Println("Grade D")
	} else if N >= 0 {
		fmt.Println("Grade E")
	} else if N < 0 {
		fmt.Println("Nilai Invalid")
	}
}
