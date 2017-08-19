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
	Shibor      = "Shibor"
	ShiborQuote = "ShiborQuote"
	ShiborMa    = "ShiborMa"
	Lpr         = "Lpr"
	LprMa       = "LprMa"
)

type BankAPI struct {
	Config Configuration
}

func NewBankAPI() *BankAPI {
	config := NewConfig()
	return &BankAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api BankAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// GetShibor 获取shibor拆放利率
	case Shibor:
		params.Add("year", c.PostForm("year"))
		api.Config.UriPath = library.UrlBankShibor
		api.Config.DataStruct = &resource.Shibor{}
	// 获取银行报价数据
	case ShiborQuote:
		params.Add("year", c.PostForm("year"))
		api.Config.UriPath = library.UrlBankShiborQuote
		api.Config.DataStruct = &resource.ShiborQuote{}
	// 获取shibor均值数据
	case ShiborMa:
		params.Add("year", c.PostForm("year"))
		api.Config.UriPath = library.UrlBankShiborMa
		api.Config.DataStruct = &resource.ShiborMa{}
	// 获取贷款基础利率LPR
	case Lpr:
		params.Add("year", c.PostForm("year"))
		api.Config.UriPath = library.UrlBankLpr
		api.Config.DataStruct = &resource.Lpr{}
	// 获取贷款基础利率均值
	case LprMa:
		params.Add("year", c.PostForm("year"))
		api.Config.UriPath = library.UrlBankLprMa
		api.Config.DataStruct = &resource.LprMa{}
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
