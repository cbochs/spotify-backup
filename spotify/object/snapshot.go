package object

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

// SnapshotID
type SnapshotID struct {
	ID       string
	Revision int
	Hash     string
}

// MarshalJSON implements serialization for SnapshotID
func (s SnapshotID) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ID)
}

// UnmarshalJSON implements deserialization for SnapshotID
func (s *SnapshotID) UnmarshalJSON(b []byte) error {
	var encoded string
	if err := json.Unmarshal(b, &encoded); err != nil {
		return err
	}

	combined, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	split := strings.Split(string(combined), ",")
	s.ID = encoded
	s.Revision, _ = strconv.Atoi(split[0])
	s.Hash = split[1]

	return nil
}
