package types

import (
	"encoding/json"
	"time"
)

// Timestamp wrapper to be able to unmarshal unix time
type Timestamp struct {
	time.Time
}

// NewTimestamp - created new Timestamp based on time.Time
func NewTimestamp(t time.Time) Timestamp {
	return Timestamp{t}
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	// 1. Decode the bytes into an int64
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		return err
	}

	// 2 - Parse the unix timestamp
	*&p.Time = time.Unix(raw, 0)
	return nil
}
