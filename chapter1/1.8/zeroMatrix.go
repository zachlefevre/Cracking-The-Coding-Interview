/*Write an algorithm that, when given an M * N matrix, zeroes out the entire column and row of any
cell which contains 0.

Ideas:
	Another data structure MUST be saved, because editing the matrix on the fly will end up
		zeroing out the entire matrix. The information regarding the col and row to zero out
		must be kept outside of the matrix.

	A map can hold the information about what cells are 0 and what cells are 1 (unchanged).
	A bitvector can hold that same information.

	If no other data structure can be used, it would be possible to iterate through the matrix
		l -> r, t -> b, and when a 0 is found in cell [a][b] to iterate through
		the section of the array already traversed, zeroing any cell in row a or col b.


*/

package main

import (
	"fmt"
)

func printMatrix(m [][]int) {
	for _, i := range m {
		fmt.Println(i)
	}
}

func zeroMatrix(m [][]int) {
	z := make(map[[2]int]bool)

	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 {
				z[[2]int{i, j}] = true
			}
		}
	}
	fmt.Println(z)
	for k := range z {
		row := m[k[0]]
		for r := range row {
			row[r] = 0
		}
		col := k[1]
		for row := range m {
			m[row][col] = 0
		}

	}
}

func main() {
	matrix := [][]int{[]int{5, 3, 9, 0, 4, 5}, []int{5, 0, 9, 7, 4, 5}, []int{5, 4, 9, 7, 4, 5}, []int{5, 1, 0, 7, 4, 5}, []int{5, 1, 8, 7, 4, 5}, []int{5, 1, 9, 7, 4, 0}}
	printMatrix(matrix)
	zeroMatrix(matrix)
	fmt.Println()
	printMatrix(matrix)

}
