//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package storage_test -destination $PWD/pkg/storage/mock_executor_test.go github.com/kamilsk/guard/pkg/storage/internal Executor
package storage_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/guard/pkg/storage"
)

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { Must(func(*Storage) error { return nil }) })
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { Must(func(*Storage) error { return errors.New("test") }) })
	})
}
