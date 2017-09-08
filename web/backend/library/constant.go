package library

import "time"

const (
	ErrGateway = "网关错误"
	ErrJobLock = "任务已上锁"

	URLDealK         = "/deal/getKData"
	URLDealHist      = "/deal/getHistData"
	URLDealTodayAll  = "/deal/getTodayAll"
	URLDealTodayTick = "/deal/getTodayTicks"
	URLDealTick      = "/deal/getTickData"
	URLDealIndex     = "/deal/getIndex"
	URLDealRTQuotes  = "/deal/getRealtimeQuotes"
	URLDealSinaDD    = "/deal/getSinaDD"

	URLBankShibor      = "/bank/getShiborData"
	URLBankShiborMa    = "/bank/getShiborMaData"
	URLBankShiborQuote = "/bank/getShiborQuoteData"
	URLBankLpr         = "/bank/getLprData"
	URLBankLprMa       = "/bank/getLprMaData"

	URLBillBrokerTops = "/billboard/getBrokerTops"
	URLBillCapTops    = "/billboard/getCapTop"
	URLBillInstDetail = "/billboard/getInstDetail"
	URLBillInstTop    = "/billboard/getInstTop"
	URLBillTopList    = "/billboard/getTopList"

	URLNewsLatestNews = "/invest/getLatestNews"
	URLNewsNotices    = "/invest/getNotices"
	URLNewsGubaSina   = "/invest/getGubaSina"

	URLClassifyArea       = "/classify/getAreaClassified"
	URLClassifyConcept    = "/classify/getConceptClassified"
	URLClassifyIndustry   = "/classify/getIndustryClassified"
	URLClassifySME        = "/classify/getSMEClassified"
	URLClassifyGEM        = "/classify/getGEMClassified"
	URLClassifyST         = "/classify/getSTClassified"
	URLClassifyHS300      = "/classify/getHS300s"
	URLClassifySZ50       = "/classify/getSZ50s"
	URLClassifyZZ500      = "/classify/getZZ500s"
	URLClassifyTerminated = "/classify/getTerminated"
	URLClassifySuspended  = "/classify/getSuspended"

	URLCompanyStockBasic = "/company/getStockBasics"
	URLCompanyProfit     = "/company/getProfitData"
	URLCompanyReport     = "/company/getReportData"
	URLCompanyOperation  = "/company/getOperationData"
	URLCompanyGrowth     = "/company/getGrowthData"
	URLCompanyDebt       = "/company/getDebtpayingData"
	URLCompanyCashflow   = "/company/getCashflowData"

	URLEconomicDepositRate    = "/economic/getDepositRate"
	URLEconomicLoanRate       = "/economic/getLoanRate"
	URLEconomicRRRRate        = "/economic/getRRRRate"
	URLEconomicMoneySupply    = "/economic/getMoneySupply"
	URLEconomicMoneySupplyBal = "/economic/getMoneySupplyBal"
	URLEconomicGDPYear        = "/economic/getGDPYear"
	URLEconomicGDPQuarter     = "/economic/getGDPQuarter"
	URLEconomicGDPFor         = "/economic/getGDPFor"
	URLEconomicGDPPull        = "/economic/getGDPPull"
	URLEconomicGDPContrib     = "/economic/getGDPContrib"
	URLEconomicCPI            = "/economic/getCPI"
	URLEconomicPPI            = "/economic/getPPI"

	URLInvestProfit         = "/invest/getProfitData"
	URLInvestForecast       = "/invest/getForecastData"
	URLInvestXSG            = "/invest/getXSGData"
	URLInvestFundHolding    = "/invest/getFundHoldings"
	URLInvestNewStock       = "/invest/getNewStocks"
	URLInvestSHMargin       = "/invest/getSHMargins"
	URLInvestSHMarginDetail = "/invest/getSHMarginsDetail"
	URLInvestSZMargin       = "/invest/getSZMargins"
	URLInvestSZMarginDetail = "/invest/getSZMarginsDetails"

	URLSinaIndustryIndex = "vip.stock.finance.sina.com.cn/q/view/newSinaHy.php"
	URLSinaConceptIndex  = "money.finance.sina.com.cn/q/view/newFLJK.php?param=class"
	URLSinaIndexDetail   = "vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=1&num=1000&sort=symbol&asc=1&node=%s&symbol=&_s_r_a=page"
	URLSinaSleep         = 1000 * time.Microsecond
	URLSinaRetryNum      = 3
	URLSinaFailBreakNum  = 10
	SchemeHttp           = "http://"
	SchemeHttps          = "https://"
)
