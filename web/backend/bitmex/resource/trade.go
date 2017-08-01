package resource

import (
	"time"
)

type Trade struct {
	Timestamp time.Time `json:"timestamp,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	Side string `json:"side,omitempty"`

	Size float32 `json:"size,omitempty"`

	Price float64 `json:"price,omitempty"`

	TickDirection string `json:"tickDirection,omitempty"`

	TrdMatchID string `json:"trdMatchID,omitempty"`

	GrossValue float32 `json:"grossValue,omitempty"`

	HomeNotional float64 `json:"homeNotional,omitempty"`

	ForeignNotional float64 `json:"foreignNotional,omitempty"`
}
