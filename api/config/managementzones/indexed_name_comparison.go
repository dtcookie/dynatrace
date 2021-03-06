package managementzones

import "github.com/dtcookie/gojson"

// IndexedNameComparison Comparison for `INDEXED_NAME` attributes.
type IndexedNameComparison struct {
	BaseComparisonBasic `json:",type=INDEXED_NAME"`
	Operator            IndexedNameComparisonOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value               *string                       `json:"value,omitempty"` // The value to compare to.
}

// UnmarshalJSON provides custom JSON deserialization
func (inc *IndexedNameComparison) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, inc)
}

// MarshalJSON provides custom JSON serialization
func (inc *IndexedNameComparison) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(inc)
}
