package managementzones

import "github.com/dtcookie/gojson"

// SimpleHostTech The value to compare to.
type SimpleHostTech struct {
	Type         *SimpleHostTechType `json:"type,omitempty"`         // Predefined technology, if technology is not predefined, then the verbatim type must be set
	VerbatimType *string             `json:"verbatimType,omitempty"` // Non-predefined technology, use for custom technologies.
}

// UnmarshalJSON provides custom JSON deserialization
func (sht *SimpleHostTech) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, sht)
}

// MarshalJSON provides custom JSON serialization
func (sht *SimpleHostTech) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(sht)
}
