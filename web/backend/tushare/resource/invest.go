package resource

import "assets/web/backend/library"

type Forecast struct {
	Code       string       `json:"code"`
	Name       string       `json:"name"`
	PreEps     float32      `json:"pre_eps"`
	ReportDate library.Date `json:"report_date"`
	Type       string       `json:"typetype"`
	Range      string       `json:"range"`
}

type XSG struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	Date  library.Date `json:"date"`
	Count string       `json:"count"`
	Ratio string       `json:"ratio"`
}

type FundHolding struct {
	Code   string       `json:"code"`
	Name   string       `json:"name"`
	Clast  string       `json:"clast"`
	Count  string       `json:"count"`
	Nums   string       `json:"nums"`
	Nlast  string       `json:"nlast"`
	Amount string       `json:"amount"`
	Ratio  string       `json:"ratio"`
	Date   library.Date `json:"date"`
}

type InvestProfit struct {
	Code       string       `json:"code"`
	Shares     float32      `json:"shares"`
	Name       string       `json:"name"`
	Divi       float32      `json:"divi"`
	ReportDate library.Date `json:"report_date"`
	Year       int          `json:"year"`
}

type SHMargin struct {
	Rzye     int64        `json:"rzye"`
	Rzrqjyzl int64        `json:"rzrqjyzl"`
	Rzmre    int64        `json:"rzmre"`
	OpDate   library.Date `json:"op_date"`
	Rqmcl    int64        `json:"rqmcl"`
	Rqyl     int64        `json:"rqyl"`
	Rqylje   int64        `json:"rqylje"`
}

type SHMarginDetail struct {
	Rzye         int64        `json:"rzye"`
	StockCode    string       `json:"stockCode"`
	OpDate       library.Date `json:"op_date"`
	Rqmcl        int64        `json:"rqmcl"`
	Rzche        int64        `json:"rzche"`
	Rzmre        int64        `json:"rzmre"`
	Rqyl         int64        `json:"rqyl"`
	SecurityAbbr string       `json:"securityAbbr"`
	Rqchl        int64        `json:"rqchl"`
}

type SZMargin struct {
	Rzye   int64        `json:"rzye"`
	Rzrqye int64        `json:"rzrqye"`
	Rzmre  int64        `json:"rzmre"`
	OpDate library.Date `json:"op_date"`
	Rqmcl  int64        `json:"rqmcl"`
	Rqyl   int64        `json:"rqyl"`
	Rqye   int64        `json:"rqye"`
}

type SZMarginDetail struct {
	Rzye         int64        `json:"rzye"`
	StockCode    string       `json:"stockCode"`
	Rqmcl        int64        `json:"rqmcl"`
	OpDate       library.Date `json:"op_date"`
	Rzmre        int64        `json:"rzmre"`
	Rqyl         int64        `json:"rqyl"`
	SecurityAbbr string       `json:"securityAbbr"`
	Rqye         int64        `json:"rqye"`
	Rzrqye       int64        `json:"rzrqye"`
}

type NewStock struct {
	Code      string       `json:"code"`
	Name      string       `json:"name"`
	IssueDate string       `json:"issue_date"`
	Pe        float32      `json:"pe"`
	Funds     float32      `json:"funds"`
	Ballot    float32      `json:"ballot"`
	Price     float32      `json:"price"`
	Xcode     string       `json:"xcode"`
	Limit     float32      `json:"limit"`
	Amount    int          `json:"amount"`
	IpoDate   library.Date `json:"ipo_date"`
	Markets   int          `json:"markets"`
}
