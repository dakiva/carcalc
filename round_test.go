package carcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	// given
	value := 101.545

	// when
	rounded := Round(value, -1)
	rounded0 := Round(value, 0)
	rounded1 := Round(value, 1)
	rounded2 := Round(value, 2)

	// then
	assert.Equal(t, 101.545, rounded)
	assert.Equal(t, 102.0, rounded0)
	assert.Equal(t, 101.5, rounded1)
	assert.Equal(t, 101.55, rounded2)
}
