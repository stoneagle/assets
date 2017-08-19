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
	BrokerTops = "BrokerTops"
	CapTops    = "CapTops"
	InstDetail = "InstDetail"
	InstTops   = "InstTops"
	TopList    = "TopList"
)

type BillBoardAPI struct {
	Config Configuration
}

func NewBillBoardAPI() *BillBoardAPI {
	config := NewConfig()
	return &BillBoardAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api BillBoardAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取营业部上榜统计
	case BrokerTops:
		params.Add("days", c.PostForm("days"))
		api.Config.UriPath = library.UrlBillBrokerTops
		api.Config.DataStruct = &resource.BrokerTop{}
	// 获取个股上榜统计
	case CapTops:
		params.Add("days", c.PostForm("days"))
		api.Config.UriPath = library.UrlBillCapTops
		api.Config.DataStruct = &resource.CapTop{}
	// 获取机构成交明细
	case InstDetail:
		api.Config.UriPath = library.UrlBillInstDetail
		api.Config.DataStruct = &resource.InstDetail{}
	// 获取机构席位追踪
	case InstTops:
		params.Add("days", c.PostForm("days"))
		api.Config.UriPath = library.UrlBillInstTop
		api.Config.DataStruct = &resource.InstTop{}
	// 获取龙虎榜列表
	case TopList:
		params.Add("date", c.PostForm("date"))
		api.Config.UriPath = library.UrlBillTopList
		api.Config.DataStruct = &resource.TopList{}
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
