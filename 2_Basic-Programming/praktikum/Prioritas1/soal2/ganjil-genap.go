package main

import "fmt"

func main() {
	var a int16

	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&a)
	if a%2 != 0 {
		fmt.Println(a, " Adalah biangan Ganjil")
	} else {
		fmt.Println(a, " Adalah bilangan Genap")
	}
}
