package resource

import (
	"time"
)

type User struct {
	Id float32 `json:"id,omitempty"`

	OwnerId float32 `json:"ownerId,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	Created time.Time `json:"created,omitempty"`

	LastUpdated time.Time `json:"lastUpdated,omitempty"`

	Preferences UserPreferences `json:"preferences,omitempty"`

	TFAEnabled string `json:"TFAEnabled,omitempty"`

	AffiliateID string `json:"affiliateID,omitempty"`

	PgpPubKey string `json:"pgpPubKey,omitempty"`

	Country string `json:"country,omitempty"`
}
