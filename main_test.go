// +build !ctl

package main

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/kamilsk/guard/cmd"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/service/guard"
	"github.com/kamilsk/guard/pkg/storage"
	"github.com/kamilsk/guard/pkg/transport/grpc"
	"github.com/kamilsk/guard/pkg/transport/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// invariant
var (
	_ guard.Storage         = storage.Must()
	_ grpc.ProtectedStorage = storage.Must()
	_ grpc.Maintenance      = guard.New(config.ServiceConfig{}, nil)
	_ http.Service          = guard.New(config.ServiceConfig{}, nil)
)

func TestService(t *testing.T) {
	build := func() *commanderMock {
		dummy := &commanderMock{}
		dummy.On("AddCommand", cmd.Completion, cmd.Migrate, cmd.Run, cmd.Version)
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
			service(tc.executor, ioutil.Discard, shutdown)
		})
	}
}
