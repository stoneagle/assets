package resource

type Leaderboard struct {
	Name string `json:"name,omitempty"`

	IsRealName bool `json:"isRealName,omitempty"`

	IsMe bool `json:"isMe,omitempty"`

	Profit float64 `json:"profit,omitempty"`
}
