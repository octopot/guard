package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Units of rate limits.
const (
	// RPS - requests per second.
	RPS RateUnit = "rps"
	// RPM - requests per minute.
	RPM RateUnit = "rpm"
	// RPH - requests per hour.
	RPH RateUnit = "rph"
	// RPD - requests per day.
	RPD RateUnit = "rpd"
	// RPW - requests per week.
	RPW RateUnit = "rpw"
)

var rate = regexp.MustCompile(`(?i:(\d+) (rps|rpm|rph|rpd|rpw)$)`)

// PackRate TODO issue#docs
func PackRate(value RateValue, unit RateUnit) Rate {
	return Rate(fmt.Sprintf("%d %s", value, unit))
}

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
func (r Rate) Value() (uint32, string) {
	if !r.IsValid() {
		return 0, ""
	}
	matches := rate.FindStringSubmatch(string(r))
	val, _ := strconv.Atoi(matches[1])
	return uint32(val), strings.ToLower(matches[2])
}

// RateValue TODO issue#docs
type RateValue uint32

// RateUnit TODO issue#docs
type RateUnit string
