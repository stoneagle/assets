package model

// HistoryData 对象
type HistoryData struct {
	Msg    string    `json:"s,omitempty"`
	Time   []float64 `json:"t,omitempty"`
	Close  []float64 `json:"c,omitempty"`
	Open   []float64 `json:"o,omitempty"`
	Low    []float64 `json:"l,omitempty"`
	High   []float64 `json:"h,omitempty"`
	Volumn []float64 `json:"v,omitempty"`
}

// HistoryModel 在influxdb中的存储对象结构
type HistoryModel struct {
	time   float64
	close  float64
	open   float64
	low    float64
	high   float64
	volumn float64
}

// HistoryModelTag 在influxdb中的tag结构
type HistoryModelTag struct {
	symbol string
}

const (
	// HistoryDayTable 日间隔数据
	HistoryDayTable = "historyDay"
	// HistoryFiveMinTable 5分钟间隔数据
	HistoryFiveMinTable = "historyFiveMin"

	// HistoryDayResolution 日间隔标记
	HistoryDayResolution = "D"
	// HistoryFiveMinResolution 5分钟间隔数据
	HistoryFiveMinResolution = "5"

	// BitmexBaseURI 爬虫基础链接
	BitmexBaseURI = "https://www.bitmex.com/api/udf/history"
)
