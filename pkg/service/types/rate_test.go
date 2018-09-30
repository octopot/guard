package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/guard/pkg/service/types"
)

func TestRate(t *testing.T) {
	type entity struct {
		Rate
		measure uint32
		unit    string
	}

	tests := []struct {
		name   string
		entity entity
	}{
		{"rate is empty", entity{"", 0, ""}},
		{"rate is invalid", entity{"abc-def-ghi", 0, ""}},
		{"rate in lowercase", entity{"10 rpm", 10, "rpm"}},
		{"rate in uppercase", entity{"10 RPM", 10, "rpm"}},
	}

	for _, test := range tests {
		measure, unit := test.entity.Value()
		assert.Equal(t, test.entity.measure, measure)
		assert.Equal(t, test.entity.unit, unit)
		assert.Equal(t, test.entity.Rate, Rate(test.entity.String()), test.name)
	}
}
