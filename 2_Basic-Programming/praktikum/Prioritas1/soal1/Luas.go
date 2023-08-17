package main

import "fmt"

func main() {
	var a float32
	var b float32
	var t float32
	var L float32

	fmt.Print("Masukkan sisi pertama = ")
	fmt.Scan(&a)
	fmt.Print("Masukkan Sisi kedua = ")
	fmt.Scan(&b)
	fmt.Print("Masukkan Tinggi Trapesium = ")
	fmt.Scan(&t)
	L = 0.5 * (a + b) * t
	fmt.Print("Luas Trapesium adalah = ", L)
}
