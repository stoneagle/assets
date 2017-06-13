package resource

import (
	"time"
)

type Quote struct {
	Timestamp time.Time `json:"timestamp,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	BidSize float32 `json:"bidSize,omitempty"`

	BidPrice float64 `json:"bidPrice,omitempty"`

	AskPrice float64 `json:"askPrice,omitempty"`

	AskSize float32 `json:"askSize,omitempty"`
}
