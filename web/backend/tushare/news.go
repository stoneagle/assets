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
	LatestNews = "LatestNews"
	Notices    = "Notices"
	SinaGuba   = "SinaGuba"
)

type NewsAPI struct {
	Config Configuration
}

func NewNewsAPI() *NewsAPI {
	config := NewConfig()
	return &NewsAPI{
		Config: *config,
	}
}

// GateWay gin的gate请求控制器
func (api NewsAPI) GateWay(c *gin.Context) {
	gateway := c.Param("gateway")
	params := url.Values{}
	var err error
	result := &ResponseData{}

	switch gateway {
	// 获取即时新闻
	case LatestNews:
		params.Add("top", c.PostForm("top"))
		params.Add("content", c.DefaultPostForm("content", "false"))
		api.Config.UriPath = library.URLNewsLatestNews
		api.Config.DataStruct = &resource.LatestNew{}
	// 获取个股信息地雷
	case Notices:
		params.Add("code", c.PostForm("code"))
		params.Add("date", c.PostForm("date"))
		api.Config.UriPath = library.URLNewsNotices
		api.Config.DataStruct = &resource.Notice{}
	// 获取sina股吧重点消息
	case SinaGuba:
		params.Add("content", c.DefaultPostForm("content", "false"))
		api.Config.UriPath = library.URLNewsGubaSina
		api.Config.DataStruct = &resource.SinaGuba{}
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
