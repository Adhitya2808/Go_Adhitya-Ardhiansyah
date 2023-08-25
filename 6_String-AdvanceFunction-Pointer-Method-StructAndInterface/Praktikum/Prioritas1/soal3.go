package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var hasil string
	length := len(a)
	if len(b) < len(a) {
		length = len(b)
	}

	for i := 0; i < length; i++ {
		if a[i] == b[i] {
			hasil += string(a[i])
 		} else{
			break
		}
	}
	return hasil
}
func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}