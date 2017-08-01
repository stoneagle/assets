package resource

import (
	"time"
)

// APIKey 对象
type APIKey struct {
	ID string `json:"id,omitempty"`

	Secret string `json:"secret,omitempty"`

	Name string `json:"name,omitempty"`

	Nonce float32 `json:"nonce,omitempty"`

	Cidr string `json:"cidr,omitempty"`

	Permissions []XAny `json:"permissions,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	UserId float32 `json:"userId,omitempty"`

	Created time.Time `json:"created,omitempty"`
}
