package carcalc

// CostInfo stores basic car cost and reductions
type CostInfo struct {
	MSRP            float64
	DownPayment     float64
	Reductions      float64
	RegistrationFee float64
	DocumentFee     float64
	TaxRate         float64
}

// Valid returns true if the info struct is in a state that will yield appropriate computed results
func (c *CostInfo) Valid() bool {
	return c.MSRP > 0 &&
		c.TaxRate > 0 && c.TaxRate < 1
}

// AdjustedCost returns the fully adjusted cost taking into account any reductions
// and fees. For a lease this value represents  the adjusted capitalized cost.
// For financing this value represents the pre-tax finance amount.
func (c *CostInfo) AdjustedCost() float64 {
	if !c.Valid() {
		return -1
	}
	return c.MSRP + c.RegistrationFee + c.DocumentFee - c.DownPayment - c.Reductions
}
