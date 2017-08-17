package tushare

import (
	"net/url"

	"assets/web/backend/library"
	"assets/web/backend/tushare/resource"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

const (
	KData          = "KData"
	HistData       = "HistData"
	TodayAll       = "TodayAll"
	TodayTick      = "TodayTick"
	Index          = "Index"
	SinaDD         = "SinaDD"
	TickData       = "TickData"
	RealTimeQuotes = "RTQuotes"
)

type DealAPI struct {
	Config Configuration
}

func NewDealAPI() *DealAPI {
	config := NewConfig()
	return &DealAPI{
		Config: *config,
	}
}

// 获取K线数据
func (d DealAPI) GetKData() (result *ResponseData, err error) {
	return
}

// 获取历史数据
func (d DealAPI) GetHistData(params url.Values, uri string) (result *ResponseData, err error) {
	return
}

// 获取当前交易所有股票的行情数据（如果是节假日，即为上一交易日
func getTodayAll(c *gin.Context) {

}

// 获取个股以往交易历史的分笔数据明细
func getTickData(c *gin.Context) {

}

// 获取实时分笔数据，可以实时取得股票当前报价和成交信息
func getRealTimeQuotes(c *gin.Context) {

}

// 获取当前交易日（交易进行中使用）已经产生的分笔明细
func getTodayTick(c *gin.Context) {

}

// 获取大盘指数行情列表
func getIndex(c *gin.Context) {

}

// 获取大单交易数据，默认为大于等于400手
func getSinaDD(c *gin.Context) {

}

// GateWay gin的gate请求控制器
func (d DealAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	case KData:
		params.Add("code", c.PostForm("code"))
		params.Add("start", c.PostForm("start"))
		params.Add("end", c.PostForm("end"))
		params.Add("ktype", c.DefaultPostForm("ktype", "D"))
		params.Add("autype", c.DefaultPostForm("autype", "qfq"))
		d.Config.UriPath = "/deal/getKData"
		d.Config.DataStruct = &resource.KData{}
	case HistData:
		params.Add("code", c.PostForm("code"))
		params.Add("start", c.PostForm("start"))
		params.Add("end", c.PostForm("end"))
		params.Add("ktype", c.DefaultPostForm("ktype", "D"))
		d.Config.UriPath = "/deal/getHistData"
		d.Config.DataStruct = &resource.HistoryData{}
	case TodayAll:
		// TODO，数据太大，返回错误
		params.Add("code", c.PostForm("code"))
		params.Add("date", c.PostForm("date"))
		d.Config.UriPath = "/deal/getTodayAll"
		d.Config.DataStruct = &resource.TodayAll{}
	case TodayTick:
		params.Add("code", c.PostForm("code"))
		d.Config.UriPath = "/deal/getTodayTicks"
		d.Config.DataStruct = &resource.TodayTick{}
	case Index:
		d.Config.UriPath = "/deal/getIndex"
		d.Config.DataStruct = &resource.Index{}
	case SinaDD:
		params.Add("code", c.PostForm("code"))
		params.Add("date", c.PostForm("date"))
		params.Add("vol", c.DefaultPostForm("vol", "400"))
		d.Config.UriPath = "/deal/getSinaDD"
		d.Config.DataStruct = &resource.SinaDD{}
	case TickData:
		params.Add("code", c.PostForm("code"))
		params.Add("date", c.PostForm("date"))
		d.Config.UriPath = "/deal/getTickData"
		d.Config.DataStruct = &resource.TickData{}
	case RealTimeQuotes:
		if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
			seelog.Errorf("获取数组参数失败，err : %+v", err)
			library.OutputErr(c, err, 401)
		}
		for _, v := range c.Request.MultipartForm.Value["symbols"] {
			params.Add("symbols", v)
		}
		d.Config.UriPath = "/deal/getRealtimeQuotes"
		d.Config.DataStruct = &resource.RealTimeQuote{}
	default:
		seelog.Errorf(library.ErrGateway)
		library.OutputErr(c, err, 401)
		return
	}

	d.Config.Params = params
	result, err = d.Config.doHttp()

	if err != nil {
		seelog.Errorf("请求失败，err = %+v\n", err)
		library.OutputErr(c, err, 401)
		return
	}

	data := d.Config.DataStruct
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

	// 动态调用方案有几个问题，1无法跳转，2无法获取结构并使用
	// result, err := lib.CallFuncByName(funcMap, gateway, d, params, uriMap[gateway])
}
