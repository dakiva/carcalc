package carcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInvalidLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidResidualLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: 1.58,
		MoneyFactor:  .0056,
		Term:         36,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidMoneyFactorLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: .58,
		MoneyFactor:  1.0056,
		Term:         36,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestInvalidTermLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: .58,
		MoneyFactor:  .0056,
		Term:         4,
	}

	// when
	valid := info.Valid()

	// then
	assert.False(t, valid)
}

func TestValidLeaseInfo(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            19000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: .58,
		MoneyFactor:  .0056,
		Term:         36,
	}

	// when
	valid := info.Valid()

	// then
	assert.True(t, valid)
}

func TestResidualValue(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: .6,
		MoneyFactor:  .0056,
		Term:         36,
	}

	// when
	residualValue := info.ResidualValue()

	// then
	assert.Equal(t, 12000.0, residualValue)
}

func TestGrossCapitalizedCost(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			TaxRate:         .0775,
		},
		ResidualRate: .6,
		MoneyFactor:  .0056,
		Term:         36,
	}

	// when
	cost := info.GrossCapitalizedCost()

	// then
	assert.Equal(t, 20180.0, cost)
}

func TestDepreciationAmount(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			DownPayment:     1000,
			TaxRate:         .0775,
		},
		ResidualRate: .59,
		MoneyFactor:  .0056,
		Term:         36,
	}

	// when
	amount := info.DepreciationAmount()

	// then
	assert.Equal(t, 7380.0, amount)
}

func TestBasePayment(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			DownPayment:     1000,
			TaxRate:         .0775,
		},
		ResidualRate: .59,
		MoneyFactor:  .0056,
		Term:         24,
	}

	// when
	amount := info.BasePayment()

	// then
	assert.Equal(t, 307.5, amount)
}

func TestRentCharge(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			DownPayment:     1000,
			TaxRate:         .0775,
		},
		ResidualRate: .59,
		MoneyFactor:  .0056,
		Term:         24,
	}

	// when
	rent := info.RentCharge()

	// then
	assert.Equal(t, 173.488, rent)
}

func TestPretaxPayment(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            20000,
			RegistrationFee: 100,
			DocumentFee:     80,
			DownPayment:     1000,
			TaxRate:         .0775,
		},
		ResidualRate: .59,
		MoneyFactor:  .0056,
		Term:         24,
	}

	// when
	payment := info.PretaxPayment()

	// then
	assert.Equal(t, 480.988, payment)
}

func TestLeasePayment(t *testing.T) {
	// given
	info := &LeaseInfo{
		CostInfo: CostInfo{
			MSRP:            23000,
			RegistrationFee: 1100,
			DocumentFee:     100,
			DownPayment:     1700,
			Reductions:      2500,
			TaxRate:         .095,
		},
		ResidualRate: .57,
		MoneyFactor:  .00125,
		Term:         36,
	}

	// when
	payment := Round(info.Payment(), 2)

	// then
	assert.Equal(t, 254.89, payment)
}
