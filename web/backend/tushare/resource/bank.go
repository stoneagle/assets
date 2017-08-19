package resource

type Shibor struct {
	Date int64   `json:"date"`
	ON   float32 `json:"ON"`
	W1   float32 `json:"1W"`
	W2   float32 `json:"2W"`
	M1   float32 `json:"1M"`
	M3   float32 `json:"3M"`
	M6   float32 `json:"6M"`
	M9   float32 `json:"9M"`
	Y1   float32 `json:"1Y"`
}

type ShiborQuote struct {
	Bank string  `json:"bank"`
	Date int64   `json:"date"`
	ONA  float32 `json:"ON_A"`
	M1A  float32 `json:"1M_A"`
	M3A  float32 `json:"3M_A"`
	M6A  float32 `json:"6M_A"`
	M9A  float32 `json:"9M_A"`
	W1A  float32 `json:"1W_A"`
	W2A  float32 `json:"2W_A"`
	Y1A  float32 `json:"1Y_A"`
	ONB  float32 `json:"ON_B"`
	M1B  float32 `json:"1M_B"`
	M3B  float32 `json:"3M_B"`
	M6B  float32 `json:"6M_B"`
	M9B  float32 `json:"9M_B"`
	W1B  float32 `json:"1W_B"`
	W2B  float32 `json:"2W_B"`
	Y1B  float32 `json:"1Y_B"`
}

type ShiborMa struct {
	Date int64   `json:"date"`
	ON5  float32 `json:"ON_5"`
	ON10 float32 `json:"ON_10"`
	ON20 float32 `json:"ON_20"`
	M15  float32 `json:"1M_5"`
	M110 float32 `json:"1M_10"`
	M120 float32 `json:"1M_20"`
	M35  float32 `json:"3M_5"`
	M310 float32 `json:"3M_10"`
	M320 float32 `json:"3M_20"`
	M65  float32 `json:"6M_5"`
	M610 float32 `json:"6M_10"`
	M620 float32 `json:"6M_20"`
	M95  float32 `json:"9M_5"`
	M910 float32 `json:"9M_10"`
	M920 float32 `json:"9M_20"`
	W15  float32 `json:"1W_5"`
	W110 float32 `json:"1W_10"`
	W120 float32 `json:"1W_20"`
	W25  float32 `json:"2W_5"`
	W210 float32 `json:"2W_10"`
	W220 float32 `json:"2W_20"`
	Y15  float32 `json:"1Y_5"`
	Y110 float32 `json:"1Y_10"`
	Y120 float32 `json:"1Y_20"`
}

type Lpr struct {
	Date int64   `json:"date"`
	Y1   float32 `json:"1Y"`
}

type LprMa struct {
	Date int64   `json:"date"`
	Y15  float32 `json:"1Y_5"`
	Y110 float32 `json:"1Y_10"`
	Y120 float32 `json:"1Y_20"`
}
