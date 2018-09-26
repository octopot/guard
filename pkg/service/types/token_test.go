package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/guard/pkg/service/types"
)

func TestToken(t *testing.T) {
	tests := []struct {
		name    string
		token   Token
		isValid bool
	}{
		{"Token is empty", "", false},
		{"Token is invalid", "abc-def-ghi", false},
		{"Token is not UUID v4", "41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false},
		{"Token in lowercase", "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true},
		{"Token in uppercase", "41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.isValid, test.token.IsValid(), test.name)
		assert.Equal(t, test.token, Token(test.token.String()), test.name)
	}
}
