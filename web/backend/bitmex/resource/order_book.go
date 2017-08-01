package resource

import (
	"time"
)

type OrderBook struct {
	Symbol string `json:"symbol,omitempty"`

	Level float32 `json:"level,omitempty"`

	BidSize float32 `json:"bidSize,omitempty"`

	BidPrice float64 `json:"bidPrice,omitempty"`

	AskPrice float64 `json:"askPrice,omitempty"`

	AskSize float32 `json:"askSize,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`
}
