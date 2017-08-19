package resource

import "assets/web/backend/library"

type DepositRate struct {
	DepositType string       `json:"deposit_type"`
	Date        library.Date `json:"date"`
	Rate        string       `json:"rate"`
}

type LoanRate struct {
	LoanType string       `json:"loan_type"`
	Date     library.Date `json:"date"`
	Rate     string       `json:"rate"`
}

type RRRRate struct {
	Now    string       `json:"now"`
	Before string       `json:"before"`
	Change string       `json:"change"`
	Date   library.Date `json:"date"`
}

type MoneySupply struct {
	Sd       string `json:"sd"`
	M0       string `json:"m0"`
	Month    string `json:"month"`
	Ftd      string `json:"ftd"`
	CdYoy    string `json:"cd_yoy"`
	M2Yoy    string `json:"m2_yoy"`
	RestsYoy string `json:"rests_yoy"`
	M2       string `json:"m2"`
	M1       string `json:"m1"`
	FtdYoy   string `json:"ftd_yoy"`
	M1Yoy    string `json:"m1_yoy"`
	Rests    string `json:"rests"`
	Qm       string `json:"qm"`
	Cd       string `json:"cd"`
	M0Yoy    string `json:"m0_yoy"`
	SdYoy    string `json:"sd_yoy"`
	QmYoy    string `json:"qm_yoy"`
}

type MoneySupplyBal struct {
	Year  string `json:"year"`
	Ftd   string `json:"ftd"`
	Qm    string `json:"qm"`
	Cd    string `json:"cd"`
	Rests string `json:"rests"`
	Sd    string `json:"sd"`
	M2    string `json:"m2"`
	M1    string `json:"m1"`
	M0    string `json:"m0"`
}

type GDPYear struct {
	Gnp           string `json:"gnp"`
	ConsIndustry  string `json:"cons_industry"`
	Pi            int64  `json:"pi"`
	Gdp           int64  `json:"gdp"`
	Si            int64  `json:"si"`
	PcGdp         string `json:"pc_gdp"`
	Lbdy          string `json:"lbdy"`
	Industry      string `json:"industry"`
	TransIndustry string `json:"trans_industry"`
	Year          int    `json:"year"`
	Ti            int    `json:"ti"`
}

type GDPQuarter struct {
	GdpYoy  float32 `json:"gdp_yoy"`
	Gdp     int64   `json:"gdp"`
	PiYoy   float32 `json:"pi_yoy"`
	Ti      int64   `json:"ti"`
	Si      int64   `json:"si"`
	SiYoy   float32 `json:"si_yoy"`
	Quarter float32 `json:"quarter"`
	TiYoy   float32 `json:"ti_yoy"`
	Pi      int64   `json:"pi"`
}

type GDPFor struct {
	EndFor     float32 `json:"end_for"`
	GoodsRate  float32 `json:"goods_rate"`
	ForRate    float32 `json:"for_rate"`
	AssetFor   float32 `json:"asset_for"`
	Year       int     `json:"year"`
	GoodsFor   float32 `json:"goods_for"`
	AssetsRate float32 `json:"assets_rate"`
}

type GDPPull struct {
	GdpYoy   float32 `json:"gdp_yoy"`
	Ti       float32 `json:"ti"`
	Si       float32 `json:"si"`
	Pi       float32 `json:"pi"`
	Year     int     `json:"year"`
	Industry float32 `json:"industry"`
}

type GDPContrib struct {
	GdpYoy   float32 `json:"gdp_yoy"`
	Ti       float32 `json:"ti"`
	Si       float32 `json:"si"`
	Pi       float32 `json:"pi"`
	Year     int     `json:"year"`
	Industry float32 `json:"industry"`
}

type CPI struct {
	Month string  `json:"month"`
	Cpi   float32 `json:"cpi"`
}

type PPI struct {
	Rmi      float32 `json:"rmi"`
	Ppi      float32 `json:"pp"`
	Ppiip    float32 `json:"cg"`
	Month    string  `json:"cg"`
	Cg       float32 `json:"cg"`
	Qm       float32 `json:"qm"`
	Roeu     float32 `json:"roeu"`
	Clothing float32 `json:"clothing"`
	Dcg      float32 `json:"dcg"`
	food     float32 `json:"food"`
	Pi       float32 `json:"pi"`
}
