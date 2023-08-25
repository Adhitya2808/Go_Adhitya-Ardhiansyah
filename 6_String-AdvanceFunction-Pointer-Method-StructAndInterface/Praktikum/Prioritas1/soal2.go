package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var total int
	total = 0
	for _, sc := range s.score {
		total += sc
	}
	return float64(total) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = math.MaxInt
	for i, sc := range s.score {
		if sc < min {
			min = sc
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = math.MinInt
	for i, sc := range s.score {
		if sc > max {
			max = sc
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a Student
	var name string
	fmt.Println("Input :")
	for i := 0; i < 5; i++ {
		fmt.Printf("Input %d Student’s Name :  ", i+1)
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Printf("Input %d Student Score :  ", i+1)
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println("\nOutput :")
	fmt.Println("Average Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")
}
