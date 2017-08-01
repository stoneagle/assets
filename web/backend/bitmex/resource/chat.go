package resource

import (
	"time"
)

type Chat struct {
	Id float32 `json:"id,omitempty"`

	Date time.Time `json:"date,omitempty"`

	User string `json:"user,omitempty"`

	Message string `json:"message,omitempty"`

	Html string `json:"html,omitempty"`

	FromBot bool `json:"fromBot,omitempty"`

	ChannelID float64 `json:"channelID,omitempty"`
}
