package types

import (
	"net/http"
	"strings"
)

const (
	cookieHeader    = "Cookie"
	refererHeader   = "Referer"
	userAgentHeader = "User-Agent"
	xEmployee       = "X-Employee"
	xForwardedFor   = "X-Forwarded-For"
	xIdentifier     = "X-Passport-ID"
	xLicense        = "X-License"
	xOptions        = "X-Guard-Options"
	xOriginalURI    = "X-Original-URI"
	xRealIP         = "X-Real-IP"
	xRequestID      = "X-Request-ID"
	xWorkplace      = "X-Workplace"
)

// MetadataFromRequest returns metadata from HTTP request.
func MetadataFromRequest(req *http.Request) Metadata {
	return Metadata{
		Cookies: func(cookies []*http.Cookie) map[string]string {
			converted := make(map[string]string)
			for _, cookie := range cookies {
				if cookie.HttpOnly && cookie.Secure {
					converted[cookie.Name] = cookie.Value
				}
			}
			return converted
		}(req.Cookies()),
		Headers: func(headers http.Header) map[string][]string {
			converted := make(map[string][]string)
			for key, values := range headers {
				if !strings.EqualFold(cookieHeader, key) {
					converted[key] = values
				}
			}
			return converted
		}(req.Header),
		Queries: req.URL.Query(),
	}
}

// Metadata holds request context.
type Metadata struct {
	Cookies map[string]string
	Headers map[string][]string
	Queries map[string][]string
}

// Employee TODO issue#docs
func (md Metadata) Employee() ID {
	return ID(md.Header(xEmployee))
}

// Header TODO issue#docs
func (md Metadata) Header(key string) string {
	return http.Header(md.Headers).Get(key)
}

// Identifier TODO issue#docs
func (md Metadata) Identifier() *ID {
	return ptr(md.Header(xIdentifier))
}

// License TODO issue#docs
func (md Metadata) License() ID {
	return ID(md.Header(xLicense))
}

// Workplace TODO issue#docs
func (md Metadata) Workplace() ID {
	return ID(md.Header(xWorkplace))
}

func ptr(origin string) *ID {
	if origin != "" {
		id := ID(origin)
		if id.IsValid() {
			return &id
		}
	}
	return nil
}
