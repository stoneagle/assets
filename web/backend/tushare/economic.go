package tushare

import (
	"assets/web/backend/library"
	"assets/web/backend/tushare/resource"
	"net/url"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

const (
	DepositRate    = "DepositRate"
	LoanRate       = "LoanRate"
	RRRRate        = "RRRRate"
	MoneySupply    = "MoneySupply"
	MoneySupplyBal = "MoneySupplyBal"
	GDPYear        = "GDPYear"
	GDPQuarter     = "GDPQuarter"
	GDPFor         = "GDPFor"
	GDPPull        = "GDPPull"
	GDPContrib     = "GDPContrib"
	CPI            = "CPI"
	PPI            = "PPI"
)

type EconomicAPI struct {
	Config Configuration
}

func NewEconomicAPI() *EconomicAPI {
	config := NewConfig()
	return &EconomicAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api EconomicAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取存款利率
	case DepositRate:
		api.Config.UriPath = library.UrlEconomicDepositRate
		api.Config.DataStruct = &resource.DepositRate{}
	// 获取贷款利率
	case LoanRate:
		api.Config.UriPath = library.UrlEconomicLoanRate
		api.Config.DataStruct = &resource.LoanRate{}
	// 获取存款准备金率
	case RRRRate:
		api.Config.UriPath = library.UrlEconomicRRRRate
		api.Config.DataStruct = &resource.RRRRate{}
	// 获取货币供应量
	case MoneySupply:
		api.Config.UriPath = library.UrlEconomicMoneySupply
		api.Config.DataStruct = &resource.MoneySupply{}
	// 获取货币供应量(年底余额)
	case MoneySupplyBal:
		api.Config.UriPath = library.UrlEconomicMoneySupplyBal
		api.Config.DataStruct = &resource.MoneySupplyBal{}
	// 获取国内生产总值
	case GDPYear:
		api.Config.UriPath = library.UrlEconomicGDPYear
		api.Config.DataStruct = &resource.GDPYear{}
	case GDPQuarter:
		api.Config.UriPath = library.UrlEconomicGDPQuarter
		api.Config.DataStruct = &resource.GDPQuarter{}
	// 三大需求对GDP的贡献
	case GDPFor:
		api.Config.UriPath = library.UrlEconomicGDPFor
		api.Config.DataStruct = &resource.GDPFor{}
	// 三大产业对Gdp的拉动
	case GDPPull:
		api.Config.UriPath = library.UrlEconomicGDPPull
		api.Config.DataStruct = &resource.GDPPull{}
	// 三大产业贡献率
	case GDPContrib:
		api.Config.UriPath = library.UrlEconomicGDPContrib
		api.Config.DataStruct = &resource.GDPContrib{}
	// 获取居民消费价格指数
	case CPI:
		api.Config.UriPath = library.UrlEconomicCPI
		api.Config.DataStruct = &resource.CPI{}
	// 获取工业品出厂价格指数
	case PPI:
		api.Config.UriPath = library.UrlEconomicPPI
		api.Config.DataStruct = &resource.PPI{}
	default:
		seelog.Errorf(library.ErrGateway)
		library.OutputErr(c, err, 401)
		return
	}

	api.Config.Params = params
	result, err = api.Config.doHttp()

	if err != nil {
		seelog.Errorf("请求失败，err = %+v\n", err)
		library.OutputErr(c, err, 401)
		return
	}

	data := api.Config.DataStruct
	err = mapstructure.Decode(result.Data, &data)
	if err != nil {
		seelog.Errorf("data转换失败，err = %+v\n", err)
		library.OutputErr(c, err, 401)
		return
	}

	c.JSON(200, gin.H{
		"data": data,
		"msg":  result.Errmsg,
	})
}
