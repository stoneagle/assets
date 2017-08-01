package resource

type Liquidation struct {
	OrderID string `json:"orderID,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	Side string `json:"side,omitempty"`

	Price float64 `json:"price,omitempty"`

	LeavesQty float32 `json:"leavesQty,omitempty"`
}
