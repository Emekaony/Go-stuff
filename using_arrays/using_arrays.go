package using_arrays

func Mean(nums []int) float64 {
	// fmt.Printf("The sum of these numbers is %d\n", sum(nums))
	sum := sum(nums)
	ll := len(nums)
	return float64(sum) / float64(ll)
}

func sum(nums []int) int {
	s := 0
	for _, j := range nums {
		s += j
	}
	return s
}
