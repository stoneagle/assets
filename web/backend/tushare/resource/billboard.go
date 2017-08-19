package resource

import "assets/web/backend/library"

type BrokerTop struct {
	Count   int     `json:"count"`
	Samount float32 `json:"samount"`
	Broker  string  `json:"broker"`
	Scount  int     `json:"scount"`
	Bcount  int     `json:"bcount"`
	Bamount float32 `json:"bamount"`
	Top3    string  `json:"top3"`
}

type CapTop struct {
	Count   int     `json:"count"`
	Samount float32 `json:"samount"`
	Net     float32 `json:"net"`
	Scount  int     `json:"scount"`
	Name    string  `json:"name"`
	Bcount  int     `json:"bcount"`
	Bamount float32 `json:"bamount"`
	Code    string  `json:"code"`
}

type InstDetail struct {
	Code    string       `json:"code"`
	Name    string       `json:"name"`
	Bamount float32      `json:"bamount"`
	Samount float32      `json:"samount"`
	Type    string       `json:"type"`
	Date    library.Date `json:"date"`
}

type InstTop struct {
	Samount float32 `json:"samount"`
	Scount  int     `json:"scount"`
	Net     float32 `json:"net"`
	Name    string  `json:"name"`
	Bcount  int     `json:"bcount"`
	Bamount float32 `json:"bamount"`
	Code    string  `json:"code"`
}

type TopList struct {
	Date    library.Date `json:"date"`
	Sell    string       `json:"sell"`
	Buy     string       `json:"buy"`
	Name    string       `json:"name"`
	Pchange string       `json:"pchange"`
	Bratio  string       `json:"bratio"`
	Code    string       `json:"code"`
	Reason  string       `json:"reason"`
	Amount  string       `json:"amount"`
	Sratio  string       `json:"sratio"`
}
