package carcalc

// Round rounds the value n decimal points
func Round(value float64, numDecimals int) float64 {
	if numDecimals < 0 {
		return value
	}
	incr := .5
	adj := 1.0
	for i := 0; i < numDecimals; i++ {
		incr /= 10
		adj *= 10
	}
	return float64(int((value+incr)*adj)) / adj
}
