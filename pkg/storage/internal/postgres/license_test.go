package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/guard/pkg/storage/internal"
	. "github.com/kamilsk/guard/pkg/storage/internal/postgres"
)

func TestLicenseManager(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ internal.LicenseManager = NewLicenseContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ internal.LicenseManager = NewLicenseContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ internal.LicenseManager = NewLicenseContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ internal.LicenseManager = NewLicenseContext(ctx, nil)
	})
}

func TestLicenseReader(t *testing.T) {
}
