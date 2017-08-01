package resource

import (
	"time"
)

// Announcement 对象
type Announcement struct {
	ID float32 `json:"id,omitempty"`

	Link string `json:"link,omitempty"`

	Title string `json:"title,omitempty"`

	Content string `json:"content,omitempty"`

	Date time.Time `json:"date,omitempty"`
}
