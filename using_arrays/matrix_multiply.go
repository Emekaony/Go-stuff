package using_arrays

import "fmt"

func dot(a, b []int) int {
	var result = 0
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	fmt.Printf("dot product is: %v\n", result)
	return result
}

func mat_mul(mat1, mat2 [][]int) [][]int {
	rows := len(mat1)
	cols := len(mat2[0])
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		// Each element in matrix[i] is automatically initialized to 0
	}
	fmt.Println(matrix)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// we havee to get the i_th row of A and the j_th column of B
			curr_row := mat1[i]
			fmt.Printf("Current row is: %v\n", curr_row)
			curr_column := []int{}
			for _, item := range mat2 {
				curr_column = append(curr_column, item[j])
			}
			fmt.Printf("Curr col is %v\n", curr_column)

			// calculate dot product
			dot := dot(curr_row, curr_column)
			matrix[i][j] = dot
		}
	}

	return matrix
}

func Run() {
	a := [][]int{
		{1, 2},
		{3, 4},
	}
	b := [][]int{
		{5, 6},
		{7, 8},
	}
	result := mat_mul(a, b)
	fmt.Println("result is ", result)
}
