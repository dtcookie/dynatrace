package managementzones

import "github.com/dtcookie/gojson"

// AzureSkuComparision Comparison for `AZURE_SKU` attributes.
type AzureSkuComparision struct {
	BaseComparisonBasic `json:",type=AZURE_SKU"`
	Value               *AzureSkuComparisionValue   `json:"value,omitempty"` // The value to compare to.
	Operator            AzureSkuComparisionOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

// UnmarshalJSON provides custom JSON deserialization
func (asc *AzureSkuComparision) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, asc)
}

// MarshalJSON provides custom JSON serialization
func (asc *AzureSkuComparision) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(asc)
}
