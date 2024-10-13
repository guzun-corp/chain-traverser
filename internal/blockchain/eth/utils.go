package eth

import (
	"time"
)

// normalizedBlockTime returns the Unix timestamp for 00:00:00 UTC of the day
// that is 'subtractDays' days before the given block time.
func normalizedBlockTime(blockTime uint64, subtractDays int) int64 {
	// Ensure subtractDays is at least 1
	if subtractDays < 1 {
		subtractDays = 1
	}

	// Convert block time to UTC
	blockDatetime := time.Unix(int64(blockTime), 0).UTC()

	// Get the start of the day for the block time
	startOfDay := time.Date(
		blockDatetime.Year(),
		blockDatetime.Month(),
		blockDatetime.Day(),
		0, 0, 0, 0,
		time.UTC,
	)

	// Subtract the specified number of days
	normalizedTime := startOfDay.AddDate(0, 0, -subtractDays)

	return normalizedTime.Unix()
}
