package types_test

import (
	"testing"

	"github.com/kamilsk/guard/pkg/service/types"
	"github.com/stretchr/testify/assert"
)

func TestRate(t *testing.T) {
	tests := []struct {
		name    string
		rate    types.Rate
		measure uint
		unit    string
	}{
		{"rate is empty", "", 0, ""},
		{"rate is invalid", "abc-def-ghi", 0, ""},
		{"rate in lowercase", "10 rpm", 10, "rpm"},
		{"rate in uppercase", "10 RPM", 10, "rpm"},
	}

	for _, test := range tests {
		measure, unit := test.rate.Value()
		assert.Equal(t, test.measure, measure)
		assert.Equal(t, test.unit, unit)
		assert.Equal(t, test.rate, types.Rate(test.rate.String()), test.name)
	}
}
