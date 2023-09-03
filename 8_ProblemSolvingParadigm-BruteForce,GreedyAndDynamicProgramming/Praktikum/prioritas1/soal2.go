package main

import "fmt"

func getRows(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := getRows(numRows - 1)

	row := make([]int, numRows)
	row[0] = 1
	row[numRows-1] = 1

	for i := 1; i < numRows-1; i++ {
		row[i] = triangle[numRows-2][i-1] + triangle[numRows-2][i]
	}

	triangle = append(triangle, row)

	return triangle
}

func main() {
	result := getRows(5)
	fmt.Println(result)
}
