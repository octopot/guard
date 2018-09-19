package types

import (
	"regexp"
	"strconv"
	"strings"
)

var rate = regexp.MustCompile(`(?i:(\d+) (rps|rpm|rph)$)`)

// Rate TODO issue#docs
type Rate string

// IsEmpty returns true if the rate is empty.
func (r Rate) IsEmpty() bool {
	return r == ""
}

// IsValid returns true if the rate is not empty and satisfies the regular expression.
func (r Rate) IsValid() bool {
	return !r.IsEmpty() && rate.MatchString(string(r))
}

// String implements built-in `fmt.Stringer` interface and returns the underlying string value.
func (r Rate) String() string {
	return string(r)
}

// Value TODO issue#docs
func (r Rate) Value() (uint, string) {
	if !r.IsValid() {
		return 0, ""
	}
	matches := rate.FindStringSubmatch(string(r))
	val, _ := strconv.Atoi(matches[1])
	return uint(val), strings.ToLower(matches[2])
}
