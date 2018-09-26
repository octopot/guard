package middleware

import (
	"context"
	"net/http"
	"strings"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/pkg/errors"
)

const (
	// AuthHeader defines authorization header.
	AuthHeader = "authorization"
	// AuthScheme defines authorization scheme.
	AuthScheme = "bearer"
)

// TokenInjector TODO issue#docs
func TokenInjector(req *http.Request) (*http.Request, error) {
	header := req.Header.Get(AuthHeader)
	if header == "" {
		return nil, errors.New("auth token not found")
	}
	splits := strings.SplitN(header, " ", 2)
	if len(splits) < 2 {
		return nil, errors.New("bad authorization string")
	}
	scheme, token := splits[0], splits[1]
	if !strings.EqualFold(scheme, AuthScheme) {
		return nil, errors.Errorf("request unauthenticated with %s", AuthScheme)
	}
	value := domain.Token(token)
	if !value.IsValid() {
		return nil, errors.Errorf("invalid auth token: %s", token)
	}
	return req.WithContext(context.WithValue(req.Context(), tokenKey{}, value)), nil
}

// TokenExtractor TODO issue#docs
func TokenExtractor(req *http.Request) (domain.ID, error) {
	value, found := req.Context().Value(tokenKey{}).(domain.ID)
	if !found {
		return value, errors.New("auth token not found")
	}
	return value, nil
}
