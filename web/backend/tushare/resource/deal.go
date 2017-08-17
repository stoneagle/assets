package resource

import (
	lib "assets/web/backend/library"
)

type HistoryData struct {
	Close       float32  `json:"close"`
	Vma10       float32  `json:"v_ma10"`
	Vma20       float32  `json:"v_ma20"`
	Low         float32  `json:"low"`
	Volume      float32  `json:"volume"`
	Ma10        float32  `json:"ma10"`
	Pchange     float32  `json:"p_change"`
	High        float32  `json:"high"`
	Turnover    float32  `json:"turnover"`
	Ma5         float32  `json:"ma5"`
	Open        float32  `json:"open"`
	Date        lib.Date `json:"date"`
	PriceChange float32  `json:"price_change"`
	Vma5        float32  `json:"v_ma5"`
	Ma20        float32  `json:"ma20"`
}

type KData struct {
	Open   float32  `json:"open"`
	Low    float32  `json:"low"`
	High   float32  `json:"high"`
	Close  float32  `json:"close"`
	Date   lib.Date `json:"date"`
	Code   string   `json:"code"`
	Volume float32  `json:"volume"`
}

type TickData struct {
	Volume float32  `json:"volume"`
	Time   lib.Time `json:"time"`
	Type   string   `json:"type"`
	Price  float32  `json:"price"`
	Amount float32  `json:"amount"`
	Change string   `json:"change"`
}

type TodayTick struct {
	Volume  float32  `json:"volume"`
	Type    string   `json:"type"`
	Amount  float32  `json:"amount"`
	Pchange string   `json:"pchange"`
	Time    lib.Time `json:"time"`
	Change  float32  `json:"change"`
	Price   float32  `json:"price"`
}

type SinaDD struct {
}

type TodayAll struct {
	Turnoverratio float32 `json:"turnoverratio"`
	High          float32 `json:"high"`
	Nmc           float32 `json:"nmc"`
	Per           float32 `json:"per"`
	Trade         float32 `json:"trade"`
	Code          string  `json:"code"`
	Volume        float32 `json:"volume"`
	Amount        float32 `json:"amount"`
	Pb            float32 `json:"pb"`
	Low           float32 `json:"low"`
	Mktcap        float32 `json:"mktcap"`
	Open          float32 `json:"open"`
	Name          string  `json:"name"`
	Changepercent float32 `json:"changepercent"`
	Settlement    float32 `json:"settlement"`
}

type RealTimeQuote struct {
	PreClose string   `json:"pre_close"`
	Time     lib.Time `json:"time"`
	Date     lib.Date `json:"date"`
	A1p      string   `json:"a1_p"`
	A2p      string   `json:"a2_p"`
	A3p      string   `json:"a3_p"`
	A4p      string   `json:"a4_p"`
	A5p      string   `json:"a5_p"`
	A1v      string   `json:"a1_v"`
	A2v      string   `json:"a2_v"`
	A3v      string   `json:"a3_v"`
	A4v      string   `json:"a4_v"`
	A5v      string   `json:"a5_v"`
	B1p      string   `json:"b1_p"`
	B2p      string   `json:"b2_p"`
	B3p      string   `json:"b3_p"`
	B4p      string   `json:"b4_p"`
	B5p      string   `json:"b5_p"`
	B1v      string   `json:"b1_v"`
	B2v      string   `json:"b2_v"`
	B3v      string   `json:"b3_v"`
	B4v      string   `json:"b4_v"`
	B5v      string   `json:"b5_v"`
	Ask      string   `json:"ask"`
	Price    string   `json:"price"`
	Amount   string   `json:"amount"`
	Bid      string   `json:"bid"`
	Open     string   `json:"open"`
	Volume   string   `json:"volume"`
	High     string   `json:"high"`
	Low      string   `json:"low"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
}

type Index struct {
	Amount   float32 `json:"amount"`
	Open     float32 `json:"open"`
	Low      float32 `json:"low"`
	High     float32 `json:"high"`
	Close    float32 `json:"close"`
	Preclose float32 `json:"preclose"`
	Volume   float32 `json:"volume"`
	Change   float32 `json:"change"`
	Code     string  `json:"code"`
	Name     string  `json:"name"`
}
