package postgres_test

import (
	"context"
	"testing"

	"go.octolab.org/ecosystem/guard/internal/storage/internal"
	. "go.octolab.org/ecosystem/guard/internal/storage/internal/postgres"
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
