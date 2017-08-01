package resource

import (
	"time"
)

type Wallet struct {
	Account float32 `json:"account,omitempty"`

	Currency string `json:"currency,omitempty"`

	PrevDeposited float32 `json:"prevDeposited,omitempty"`

	PrevWithdrawn float32 `json:"prevWithdrawn,omitempty"`

	PrevAmount float32 `json:"prevAmount,omitempty"`

	PrevTimestamp time.Time `json:"prevTimestamp,omitempty"`

	DeltaDeposited float32 `json:"deltaDeposited,omitempty"`

	DeltaWithdrawn float32 `json:"deltaWithdrawn,omitempty"`

	DeltaAmount float32 `json:"deltaAmount,omitempty"`

	Deposited float32 `json:"deposited,omitempty"`

	Withdrawn float32 `json:"withdrawn,omitempty"`

	Amount float32 `json:"amount,omitempty"`

	PendingCredit float32 `json:"pendingCredit,omitempty"`

	PendingDebit float32 `json:"pendingDebit,omitempty"`

	ConfirmedDebit float32 `json:"confirmedDebit,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`

	Addr string `json:"addr,omitempty"`

	WithdrawalLock []XAny `json:"withdrawalLock,omitempty"`
}
