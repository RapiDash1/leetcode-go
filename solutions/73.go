package solutions

import "math"

func setZeroes(matrix [][]int) {
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 0 {
				for iterRow := range matrix {
					if matrix[iterRow][col] != 0 {
						matrix[iterRow][col] = math.MinInt64
					}
				}
				for iterCol := range matrix[row] {
					if matrix[row][iterCol] != 0 {
						matrix[row][iterCol] = math.MinInt64
					}
				}
			}
		}
	}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == math.MinInt64 {
				matrix[row][col] = 0
			}
		}
	}
}
