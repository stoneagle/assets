package resource

type InstrumentInterval struct {
	Intervals []string `json:"intervals,omitempty"`

	Symbols []string `json:"symbols,omitempty"`
}
