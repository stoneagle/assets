package resource

type OrderBookL2 struct {
	Symbol string `json:"symbol,omitempty"`

	Id float32 `json:"id,omitempty"`

	Side string `json:"side,omitempty"`

	Size float32 `json:"size,omitempty"`

	Price float64 `json:"price,omitempty"`
}
