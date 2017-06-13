package resource

import (
	"time"
)

type Settlement struct {
	Timestamp time.Time `json:"timestamp,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	SettlementType string `json:"settlementType,omitempty"`

	SettledPrice float64 `json:"settledPrice,omitempty"`

	Bankrupt float32 `json:"bankrupt,omitempty"`

	TaxBase float32 `json:"taxBase,omitempty"`

	TaxRate float64 `json:"taxRate,omitempty"`
}
