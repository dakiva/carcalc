package carcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInvalidCostInfo(t *testing.T) {
	// given
	info := &CostInfo{}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidRateCostInfo(t *testing.T) {
	// given
	info := &CostInfo{
		MSRP:            19000,
		RegistrationFee: 100,
		DocumentFee:     80,
		TaxRate:         1.1,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestValidCostInfo(t *testing.T) {
	// given
	info := &CostInfo{
		MSRP:            19000,
		RegistrationFee: 100,
		DocumentFee:     80,
		TaxRate:         .0775,
	}

	// when
	valid := info.Valid()

	// then
	assert.True(t, valid)
}

func TestInvalidAdjustedCost(t *testing.T) {
	// given
	info := &CostInfo{
		MSRP: 19000,
	}

	// when
	cost := info.AdjustedCost()

	// then
	assert.Equal(t, -1.0, cost)
}

func TestAdjustedCost(t *testing.T) {
	// given
	info := CostInfo{
		MSRP:            20000,
		DownPayment:     1000.0,
		Reductions:      500.0,
		RegistrationFee: 200,
		DocumentFee:     50,
		TaxRate:         .0775,
	}

	// when
	cost := info.AdjustedCost()

	// then
	assert.Equal(t, 18750.0, cost)
}

func TestAdjustedCostNoReductions(t *testing.T) {
	// given
	info := CostInfo{
		MSRP:            20000,
		RegistrationFee: 200,
		DocumentFee:     50,
		TaxRate:         .0775,
	}

	// when
	cost := info.AdjustedCost()

	// then
	assert.Equal(t, 20250.0, cost)
}
