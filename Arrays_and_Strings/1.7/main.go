package main

import "fmt"

// Zero matrix
// implement an algorithm such that if an element is MXN matrix is 0
// its entire row is 0 and column are set to 0

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	setZero(matrix)

	fmt.Println(matrix)
}

func setZero(matrix [][]int) {
	lenRows := len(matrix)
	lenCols := len(matrix[0])

	row := make([]bool, lenRows)
	column := make([]bool, lenCols)

	for i := range lenRows {
		for j := range lenCols {
			if matrix[i][j] == 0 {
				row[i] = true
				column[i] = true
			}
		}
	}

	for i := range row {
		if row[i] {
			nullifyRow(matrix, i)
		}
	}

	for i := range column {
		if column[i] {
			nullifyColumn(matrix, i)
		}
	}

}

func nullifyRow(matrix [][]int, rowIdx int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[rowIdx][i] = 0
	}
}

func nullifyColumn(matrix [][]int, colIdx int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[i][colIdx] = 0
	}
}
