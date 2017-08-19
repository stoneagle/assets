package resource

type StockBasic struct {
	Name             string  `json:"name"`
	Esp              float32 `json:"esp"`
	FixedAssets      float32 `json:"fixedAssets"`
	Industry         string  `json:"industry"`
	Outstanding      float32 `json:"outstanding"`
	Holders          int     `json:"holders"`
	Area             string  `json:"area"`
	Pb               float32 `json:"pb"`
	Reserved         float32 `json:"reserved"`
	Perundp          float32 `json:"perundp"`
	LiquidAssets     float32 `json:"liquidAssets"`
	Totals           float32 `json:"totals"`
	Bvps             float32 `json:"bvps"`
	Rev              float32 `json:"rev"`
	Profit           float32 `json:"profit"`
	Npr              float32 `json:"npr"`
	Undp             float32 `json:"undp"`
	Pe               float32 `json:"pe"`
	TimeToMarket     int64   `json:"timeToMarket"`
	Gpr              float32 `json:"gpr"`
	ReservedPerShare float32 `json:"reservedPerShare"`
	TotalAssets      float32 `json:"totalAssets"`
}

type CompanyProfit struct {
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	NetProfitRatio  float32 `json:"net_profit_ratio"`
	Eps             float32 `json:"eps"`
	Roe             float32 `json:"roe"`
	BusinessIncome  float32 `json:"business_income"`
	GrossProfitRate float32 `json:"gross_profit_rate"`
	NetProfits      float32 `json:"net_profits"`
	Bips            float32 `json:"bips"`
}

type CompanyReport struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	NetProfitRatio float32 `json:"net_profit_ratio"`
	Eps            float32 `json:"eps"`
	EpsYoy         float32 `json:"eps_yoy"`
	ReportDate     string  `json:"report_date"`
	Bvps           float32 `json:"bvps"`
	Roe            float32 `json:"roe"`
	Epcf           float32 `json:"epcf"`
	Distrib        float32 `json:"distrib"`
	ProfitsYoy     float32 `json:"profits_yoy"`
}

type Operation struct {
	Code                 string  `json:"code"`
	Name                 string  `json:"name"`
	Arturnover           float32 `json:"arturnover"`
	InventoryDays        float32 `json:"inventory_days"`
	CurrentassetTurnover float32 `json:"currentasset_turnover"`
	InventoryTurnover    float32 `json:"inventory_turnover"`
	Arturndays           float32 `json:"arturndays"`
	CurrentassetDays     float32 `json:"currentasset_days"`
}

type Growth struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Seg  float32 `json:"seg"`
	Epsg float32 `json:"epsg"`
	Nav  float32 `json:"nav"`
	Nprg float32 `json:"nprg"`
	Targ float32 `json:"targ"`
	Mbrg float32 `json:"mbrg"`
}

type Cashflow struct {
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	CfLiabilities float32 `json:"cf_liabilities"`
	Rateofreturn  float32 `json:"rateofreturn"`
	CfNm          float32 `json:"cf_nm"`
	Cashflowratio float32 `json:"cashflowratio"`
	CfSales       float32 `json:"cf_sales"`
}

type DebtPaying struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Icratio      float32 `json:"icratio"`
	Currentratio string  `json:"currentratio"`
	Quickratio   string  `json:"quickratio"`
	Sheqratio    string  `json:"sheqratio"`
	Cashratio    string  `json:"cashratio"`
	Adratio      string  `json:"adratio"`
}
