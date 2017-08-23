package resource

type SinaDetail struct {
	Symbol        string  `json:"symbol"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Trade         string  `json:"trade"`
	Pricechange   string  `json:"pricechange"`
	Changepercent string  `json:"changepercent"`
	Buy           string  `json:"buy"`
	Sell          string  `json:"sell"`
	Settlement    string  `json:"settlement"`
	Open          string  `json:"open"`
	High          string  `json:"high"`
	Low           string  `json:"low"`
	Volume        int64   `json:"volume"`
	Amount        int64   `json:"amount"`
	Ticktime      string  `json:"ticktime"`
	Per           float32 `json:"per"`
	Pb            float32 `json:"pb"`
	Mktcap        float32 `json:"mktcap"`
	Nmc           float32 `json:"nmc"`
	Turnoverratio float32 `json:"turnoverratio"`
}

type SinaSimple struct {
	Symbol string `json:"symbol"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}
