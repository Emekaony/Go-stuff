package using_arrays

import "fmt"

// given two 1-D matrices, return their dot product
func dot(a, b []int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	fmt.Printf("dot product is: %v\n", result)
	return result
}

func mat_mul(mat1, mat2 [][]int) [][]int {
	rows := len(mat1)
	cols := len(mat2[0])
	// initialize the result and pad with zeros
	matrix := make([][]int, rows)
	for i := range matrix {
		// Each element in matrix[i] is automatically initialized to 0
		matrix[i] = make([]int, cols)
	}

	// the resulting matrix should have the rows of matrix1 and the columns of matrix 2
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// we havee to get the i_th row of A and the j_th column of B
			row_matrix := mat1[i]
			column_matrix := []int{}
			// iterate through each row in matrix2 and get the jth element
			// this will give us our desired column matrix
			for _, item := range mat2 {
				column_matrix = append(column_matrix, item[j])
			}
			// calculate dot product of both matrices
			dot := dot(row_matrix, column_matrix)
			matrix[i][j] = dot
		}
	}
	return matrix
}

func Run() {
	a := [][]int{
		{7, 6},
		{5, 9},
	}
	b := [][]int{
		{5, 6},
		{7, 8},
	}
	result := mat_mul(a, b)
	fmt.Println("result is ", result)
}
