package resource

import (
	"time"
)

// AccessToken 对象
type AccessToken struct {
	ID string `json:"id,omitempty"`

	// time to live in seconds (2 weeks by default)
	TTL float64 `json:"ttl,omitempty"`

	Created time.Time `json:"created,omitempty"`

	UserID float64 `json:"userId,omitempty"`
}
