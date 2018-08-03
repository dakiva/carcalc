package carcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInvalidFinanceInfo(t *testing.T) {
	// given
	info := &FinanceInfo{}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidTermFinanceInfo(t *testing.T) {
	// given
	info := &FinanceInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		Term: 4,
		APR:  .034,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidAPRFinanceInfo(t *testing.T) {
	// given
	info := &FinanceInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		Term: 12,
		APR:  1.034,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestValidFinanceInfo(t *testing.T) {
	// given
	info := &FinanceInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		Term: 12,
		APR:  .034,
	}

	// when
	valid := info.Valid()

	// then
	assert.True(t, valid)
}

func TestPrincipal(t *testing.T) {
	// given
	info := &FinanceInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		Term: 12,
		APR:  .034,
	}

	// when
	principal := Round(info.Principal(), 2)

	// then
	assert.Equal(t, 20666.45, principal)
}

func TestFinancePayment(t *testing.T) {
	// given
	info := &FinanceInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 900,
			DocumentFee:     100,
			TaxRate:         .0775,
		},
		Term: 60,
		APR:  .034,
	}

	// when
	principal := Round(info.Payment(), 2)

	// then
	assert.Equal(t, 391.07, principal)
}
