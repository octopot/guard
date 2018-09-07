package main

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

type commanderMock struct {
	mock.Mock
	commands []*cobra.Command
}

func (m *commanderMock) AddCommand(commands ...*cobra.Command) {
	m.commands = commands
	converted := make([]interface{}, 0, len(commands))
	for _, cmd := range commands {
		converted = append(converted, cmd)
	}
	m.Called(converted...)
}

func (m *commanderMock) Execute() error {
	return m.Called().Error(0)
}
