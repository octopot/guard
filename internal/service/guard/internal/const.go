package internal

import "time"

const (
	maxLicenses   = 5000
	maxWorkplaces = 25
	licenseTTL    = 15 * time.Minute
	workplaceTTL  = 24 * time.Hour
)
