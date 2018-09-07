// +build cli ctl

package main

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/kamilsk/guard/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestControl(t *testing.T) {
	tests := []struct {
		name     string
		executor commander
		expected []int
	}{
		{
			name: "success",
			executor: func() commander {
				executor := &commanderMock{}
				executor.On("AddCommand", cmd.Completion, cmd.Control, cmd.Version)
				executor.On("Execute").Return(nil)
				return executor
			}(),
			expected: []int{success},
		},
		{
			name: "failure",
			executor: func() commander {
				executor := &commanderMock{}
				executor.On("AddCommand", cmd.Completion, cmd.Control, cmd.Version)
				executor.On("Execute").Return(errors.New("test"))
				return executor
			}(),
			expected: []int{failure, success},
		},
		{
			name: "panic",
			executor: func() commander {
				executor := &commanderMock{}
				executor.On("AddCommand", cmd.Completion, cmd.Control, cmd.Version)
				executor.On("Execute").Run(func(mock.Arguments) { panic("test") })
				return executor
			}(),
			expected: []int{failure},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			shutdown := func(code int) {
				var expected int
				expected, tc.expected = tc.expected[0], tc.expected[1:]
				assert.Equal(t, expected, code)
			}
			control(tc.executor, ioutil.Discard, shutdown)
		})
	}
}
