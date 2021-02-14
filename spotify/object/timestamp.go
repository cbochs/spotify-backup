package object

import (
	"fmt"
	"time"
)

// Timestamp is a UNIX timestamp in ISO 1806 format
type Timestamp struct {
	time.Time
}

// MarshalJSON implements serialization for a timestamp
func (t Timestamp) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", t.Format(time.RFC3339Nano))
	return []byte(s), nil
}

// UnmarshalJSON implements deserialization for a timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := string(b)
	time, err := time.Parse(time.RFC3339Nano, s[1:len(s)-1])
	if err != nil {
		return err
	}

	t.Time = time

	return nil
}
