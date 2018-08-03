package carcalc

import "math"

// FinanceInfo holds financing related information
type FinanceInfo struct {
	CostInfo
	APR  float64
	Term int
}

// Valid returns true if the info struct is in a state that will yield appropriate computed results
func (f *FinanceInfo) Valid() bool {
	return f.CostInfo.Valid() &&
		f.APR > 0 && f.APR < 1 &&
		f.Term > 0 && f.Term%3 == 0
}

// Principal returns the total principal amount to be financed including sales tax
func (f *FinanceInfo) Principal() float64 {
	if !f.Valid() {
		return -1
	}
	return f.AdjustedCost() * (1 + f.TaxRate)
}

// Payment returns the final payment with all taxes included
func (f *FinanceInfo) Payment() float64 {
	if !f.Valid() {
		return -1
	}
	numerator := (f.APR / 12) * f.Principal()
	denominator := 1.0 - math.Pow(1.0+(f.APR/12.0), float64(f.Term*-1))
	return numerator / denominator
}
