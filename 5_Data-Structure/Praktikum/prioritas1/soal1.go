package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	ada := make(map[string]bool)

	gabung := make([]string, 0)

	for _, name := range arrayA {
		if !ada[name] {
			gabung = append(gabung, name)
			ada[name] = true
		}
	}

	for _, name := range arrayB {
		if !ada[name] {
			gabung = append(gabung, name)
			ada[name] = true
		}
	}

	return gabung	
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
