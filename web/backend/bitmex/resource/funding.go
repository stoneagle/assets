package resource

import (
	"time"
)

type Funding struct {
	Timestamp time.Time `json:"timestamp,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	FundingInterval time.Time `json:"fundingInterval,omitempty"`

	FundingRate float64 `json:"fundingRate,omitempty"`

	FundingRateDaily float64 `json:"fundingRateDaily,omitempty"`
}
