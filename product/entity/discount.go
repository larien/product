package entity

// Discount represents the Discount object and contains the fields that can be manipulated by the service
type Discount struct {
	Percentage   int64 `json:"percentage"`
	ValueInCents int64 `json:"value_in_cents"`
}
