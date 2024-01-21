package array_algorithms

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var primaryDiagonalSum int32
	var secondaryDiagonalSum int32
	numRows := len(arr)
	for i := 0; i < numRows; i++ {
		primaryDiagonalSum += arr[i][i]
	}
	// for initialize; boolean_expr; increment/decrement
	for i, j := 0, numRows-1; i < numRows && j >= 0; i, j = i+1, j-1 {
		secondaryDiagonalSum += arr[i][j]
	}
	diff := float64(primaryDiagonalSum - secondaryDiagonalSum)
	return int32(math.Abs(diff))
}
