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
	Area       = "Area"
	Concept    = "Concept"
	Industry   = "Industry"
	SME        = "SME"
	GEM        = "GEM"
	ST         = "ST"
	HS300      = "HS300"
	SZ50       = "SZ50"
	ZZ500      = "ZZ500"
	Suspended  = "Suspended"
	Terminated = "Terminated"
)

type ClassifyAPI struct {
	Config Configuration
}

func NewClassifyAPI() *ClassifyAPI {
	config := NewConfig()
	return &ClassifyAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api ClassifyAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取地区
	case Area:
		api.Config.UriPath = library.UrlClassifyArea
		api.Config.DataStruct = &resource.Area{}
	// 获取概念(获取数据不稳定，需要重试)
	case Concept:
		api.Config.UriPath = library.UrlClassifyConcept
		api.Config.DataStruct = &resource.Concept{}
	// 获取行业(有时获取数据不稳定，需要重试)
	case Industry:
		api.Config.UriPath = library.UrlClassifyIndustry
		api.Config.DataStruct = &resource.Industry{}
	// 获取中小板分类
	case SME:
		api.Config.UriPath = library.UrlClassifySME
		api.Config.DataStruct = &resource.SME{}
	// 获取创业板分类
	case GEM:
		api.Config.UriPath = library.UrlClassifyGEM
		api.Config.DataStruct = &resource.GEM{}
	// 获取风险警示板
	case ST:
		api.Config.UriPath = library.UrlClassifyST
		api.Config.DataStruct = &resource.ST{}
	// 获取沪深300
	case HS300:
		api.Config.UriPath = library.UrlClassifyHS300
		api.Config.DataStruct = &resource.HS300{}
	// 获取上证50
	case SZ50:
		api.Config.UriPath = library.UrlClassifySZ50
		api.Config.DataStruct = &resource.SZ50{}
	// 获取中证500
	case ZZ500:
		api.Config.UriPath = library.UrlClassifyZZ500
		api.Config.DataStruct = &resource.ZZ500{}
	// 获取暂停上市
	case Suspended:
		api.Config.UriPath = library.UrlClassifySuspended
		api.Config.DataStruct = &resource.Suspended{}
	// 获取终止上市
	case Terminated:
		api.Config.UriPath = library.UrlClassifyTerminated
		api.Config.DataStruct = &resource.Terminated{}
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
