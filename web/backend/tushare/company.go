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
	StockBasic    = "StockBasic"
	CompanyProfit = "Profit"
	CompanyReport = "Report"
	Operation     = "Operation"
	Growth        = "Growth"
	Cashflow      = "Cashflow"
	DebtPaying    = "DebtPaying"
)

type CompanyAPI struct {
	Config Configuration
}

func NewCompanyAPI() *CompanyAPI {
	config := NewConfig()
	return &CompanyAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api CompanyAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取上市公司基本情况
	case StockBasic:
		api.Config.UriPath = library.URLCompanyStockBasic
		api.Config.DataStruct = &resource.StockBasic{}
	// 获取盈利能力
	case CompanyProfit:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyProfit
		api.Config.DataStruct = &resource.CompanyProfit{}
	// 获取业绩报告
	case CompanyReport:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyReport
		api.Config.DataStruct = &resource.CompanyReport{}
	// 获取营运能力
	case Operation:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyOperation
		api.Config.DataStruct = &resource.Operation{}
	// 获取成长能力
	case Growth:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyGrowth
		api.Config.DataStruct = &resource.Growth{}
	// 获取现金流量
	case Cashflow:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyCashflow
		api.Config.DataStruct = &resource.Cashflow{}
	// 获取偿债能力
	case DebtPaying:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.URLCompanyDebt
		api.Config.DataStruct = &resource.DebtPaying{}
	default:
		seelog.Errorf(library.ErrGateway)
		library.OutputErr(c, err, 401)
		return
	}

	api.Config.Params = params
	result, err = api.Config.doPost()

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
