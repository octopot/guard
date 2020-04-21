//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package storage_test -destination mock_executor_test.go go.octolab.org/ecosystem/guard/internal/storage/internal Executor
package storage_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/guard/internal/storage"
)

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { Must(func(*Storage) error { return nil }) })
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { Must(func(*Storage) error { return errors.New("test") }) })
	})
}
