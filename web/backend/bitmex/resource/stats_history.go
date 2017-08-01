package resource

import (
	"time"
)

type StatsHistory struct {
	Date time.Time `json:"date,omitempty"`

	RootSymbol string `json:"rootSymbol,omitempty"`

	Currency string `json:"currency,omitempty"`

	Volume float32 `json:"volume,omitempty"`

	Turnover float32 `json:"turnover,omitempty"`
}
