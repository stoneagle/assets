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
	InvestProfit   = "Profit"
	Forecast       = "Forecast"
	XSG            = "XSG"
	FundHolding    = "FundHolding"
	NewStock       = "NewStock"
	SHMargin       = "SHMargin"
	SHMarginDetail = "SHMarginDetail"
	SZMargin       = "SZMargin"
	SZMarginDetail = "SZMarginDetail"
)

type InvestAPI struct {
	Config Configuration
}

func NewInvestAPI() *InvestAPI {
	config := NewConfig()
	return &InvestAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api InvestAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取业绩预告
	case Forecast:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.UrlInvestForecast
		api.Config.DataStruct = &resource.Forecast{}
	case XSG:
		params.Add("year", c.PostForm("year"))
		params.Add("month", c.PostForm("month"))
		api.Config.UriPath = library.UrlInvestXSG
		api.Config.DataStruct = &resource.XSG{}
	// 获取基金持股
	case FundHolding:
		params.Add("year", c.PostForm("year"))
		params.Add("season", c.PostForm("season"))
		api.Config.UriPath = library.UrlInvestFundHolding
		api.Config.DataStruct = &resource.FundHolding{}
	// 获取分配预案
	case InvestProfit:
		params.Add("year", c.PostForm("year"))
		params.Add("top", c.PostForm("top"))
		api.Config.UriPath = library.UrlInvestProfit
		api.Config.DataStruct = &resource.InvestProfit{}
	// 获取新股
	case NewStock:
		api.Config.UriPath = library.UrlInvestNewStock
		api.Config.DataStruct = &resource.NewStock{}
	// 获取融资融券
	case SHMargin:
		params.Add("start", c.PostForm("start"))
		params.Add("end", c.PostForm("end"))
		api.Config.UriPath = library.UrlInvestSHMargin
		api.Config.DataStruct = &resource.SHMargin{}
	case SHMarginDetail:
		params.Add("date", c.PostForm("date"))
		params.Add("symbol", c.PostForm("symbol"))
		params.Add("start", c.PostForm("start"))
		params.Add("end", c.PostForm("end"))
		api.Config.UriPath = library.UrlInvestSHMarginDetail
		api.Config.DataStruct = &resource.SHMarginDetail{}
	case SZMargin:
		params.Add("start", c.PostForm("start"))
		params.Add("end", c.PostForm("end"))
		api.Config.UriPath = library.UrlInvestSZMargin
		api.Config.DataStruct = &resource.SZMargin{}
	case SZMarginDetail:
		params.Add("date", c.PostForm("date"))
		api.Config.UriPath = library.UrlInvestSZMarginDetail
		api.Config.DataStruct = &resource.SZMarginDetail{}
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
