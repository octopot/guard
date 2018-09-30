package grpc_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/guard/pkg/transport/grpc"
)

func TestTimestamp(t *testing.T) {
	tests := []struct {
		name   string
		time   *time.Time
		assert func(assert.TestingT, *time.Time)
	}{
		{"nil pointer", nil, func(t assert.TestingT, tp *time.Time) { assert.Nil(t, Timestamp(tp)) }},
		{"normal use", new(time.Time), func(t assert.TestingT, tp *time.Time) { assert.NotNil(t, Timestamp(tp)) }},
		{"invalid time", func() *time.Time {
			tp := time.Time{}.AddDate(-math.MaxInt32, -math.MaxInt32, -math.MaxInt32)
			return &tp
		}(), func(t assert.TestingT, tp *time.Time) { assert.Panics(t, func() { Timestamp(tp) }) }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, tc.time)
		})
	}
}
