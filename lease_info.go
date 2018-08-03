package carcalc

// LeaseInfo holds lease related information
type LeaseInfo struct {
	CostInfo
	ResidualRate float64
	MoneyFactor  float64
	Term         int
}

// Valid returns true if the info struct is in a state that will yield appropriate computed results
func (l *LeaseInfo) Valid() bool {
	return l.CostInfo.Valid() &&
		l.ResidualRate > 0 && l.ResidualRate < 1 &&
		l.MoneyFactor > 0 && l.MoneyFactor < 1 &&
		l.Term > 0 && l.Term%3 == 0
}

// ResidualValue returns the residual value based on MSRP
func (l *LeaseInfo) ResidualValue() float64 {
	if !l.Valid() {
		return -1
	}
	return l.MSRP * l.ResidualRate
}

// GrossCapitalizedCost returns the capitalized cost including all fees
func (l *LeaseInfo) GrossCapitalizedCost() float64 {
	if !l.Valid() {
		return -1
	}
	return l.MSRP + l.RegistrationFee + l.DocumentFee
}

// DepreciationAmount returns the basis for the lease payment
func (l *LeaseInfo) DepreciationAmount() float64 {
	if !l.Valid() {
		return -1
	}
	return l.AdjustedCost() - l.ResidualValue()
}

// BasePayment returns the base payment
func (l *LeaseInfo) BasePayment() float64 {
	if !l.Valid() {
		return -1
	}
	return l.DepreciationAmount() / float64(l.Term)
}

// RentCharge returns the monthly rent charge
func (l *LeaseInfo) RentCharge() float64 {
	if !l.Valid() {
		return -1
	}
	return (l.ResidualValue() + l.AdjustedCost()) * l.MoneyFactor
}

// PretaxPayment returns the pretax payment including the rent charge
func (l *LeaseInfo) PretaxPayment() float64 {
	if !l.Valid() {
		return -1
	}
	return l.BasePayment() + l.RentCharge()
}

// Payment returns the final payment with all taxes included
func (l *LeaseInfo) Payment() float64 {
	if !l.Valid() {
		return -1
	}
	return l.PretaxPayment() * (1 + l.TaxRate)
}
