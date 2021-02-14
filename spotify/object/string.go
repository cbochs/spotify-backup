package object

import (
	"encoding/json"
	"html"
)

// HTMKString
type HTMLString string

// MarshalJSON implements serialization for HTMLString
func (s HTMLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(html.EscapeString(string(s)))
}

// UnmarshalJSON implements deserialization for HTMLString
func (s *HTMLString) UnmarshalJSON(b []byte) error {
	var ss string
	if err := json.Unmarshal(b, &ss); err != nil {
		*s = ""
		return err
	}
	*s = HTMLString(html.UnescapeString(ss))

	return nil
}
