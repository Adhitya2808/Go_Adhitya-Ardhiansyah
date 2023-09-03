package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(item []string) string {
	counts := make(map[string]int)

	// count the occurrences of each item
	for _, name := range item {
		counts[name]++
	}

	// add data into slice
	pairs := make([]pair, 0, len(counts))
	for name, count := range counts {
		pairs = append(pairs, pair{name, count})
	}

	n := len(pairs)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if pairs[j].count < pairs[j+1].count {
				pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			}
		}
	}


	var output string
	for i, p := range pairs {
		if i > 0 {
			output += " "
		}
		output += fmt.Sprintf("%v -> %v", p.name, p.count)
	}

	return output
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
