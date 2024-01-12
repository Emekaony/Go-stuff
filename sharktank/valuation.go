package sharktank

func Valuation(amount, percentage float64) float64 {
	// just need to multiply amount by 100/percent
	return amount * (float64(100) / percentage)
}
