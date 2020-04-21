package main

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kamilsk/guard/pkg/cmd"
)

func TestControl(t *testing.T) {
	build := func() *commanderMock {
		dummy := &commanderMock{}
		dummy.On("AddCommand", cmd.Completion, cmd.Install, cmd.License, cmd.Version)
		return dummy
	}

	tests := []struct {
		name     string
		executor commander
		expected []int
	}{
		{
			name: "success",
			executor: func() commander {
				executor := build()
				executor.On("Execute").Return(nil)
				return executor
			}(),
			expected: []int{success},
		},
		{
			name: "failure",
			executor: func() commander {
				executor := build()
				executor.On("Execute").Return(errors.New("test"))
				return executor
			}(),
			expected: []int{failure, success},
		},
		{
			name: "panic",
			executor: func() commander {
				executor := build()
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

type commanderMock struct {
	mock.Mock
	commands []*cobra.Command
}

func (m *commanderMock) AddCommand(cc ...*cobra.Command) {
	m.commands = cc
	converted := make([]interface{}, 0, len(cc))
	for _, c := range cc {
		converted = append(converted, c)
	}
	m.Called(converted...)
}

func (m *commanderMock) Execute() error {
	return m.Called().Error(0)
}
