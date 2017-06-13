package resource

import (
	"time"
)

type Insurance struct {
	Currency string `json:"currency,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`

	WalletBalance float32 `json:"walletBalance,omitempty"`
}
