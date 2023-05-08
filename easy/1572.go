package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func diagonalSum(mat [][]int) int {
	rowCount := len(mat)
	existsHash := make(map[string]bool)
	sum := 0
	anotherI := rowCount - 1
	for i := 0; i < rowCount; i++ {
		key := strconv.Itoa(i) + "-" + strconv.Itoa(i)
		if _, ok := existsHash[key]; !ok {
			sum += mat[i][i]
			existsHash[key] = true
		}

		key = strconv.Itoa(i) + "-" + strconv.Itoa(anotherI)
		if _, ok := existsHash[key]; !ok {
			sum += mat[i][anotherI]
			existsHash[key] = true
		}

		anotherI--
	}

	return sum
}
