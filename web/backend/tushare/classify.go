package tushare

import (
	"assets/web/backend/library"
	"assets/web/backend/tushare/resource"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/cihub/seelog"
	iconv "github.com/djimenez/iconv-go"
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

// 根据tag检查并更新分类详情
func checkDetailByTag(tag string, classifyName string, config Configuration) {
	host := library.SchemeHttp + library.URLSinaIndexDetail
	host = fmt.Sprintf(host, tag)
	body, err := config.doGet(host)
	if err != nil {
		seelog.Errorf("err = %+v\n", err)
		return
	}
	dataStr := string(body)
	reg := regexp.MustCompile(`\,(.*?)\:`)
	dataStr = reg.ReplaceAllString(dataStr, ",\"${1}\":")
	dataStr = strings.Replace(dataStr, "\"{symbol", "{\"symbol", -1)
	dataStr = strings.Replace(dataStr, "{symbol", "{\"symbol\"", -1)
	dataStrDecode, err := iconv.ConvertString(dataStr, "GB2312", "utf-8")
	var detail []resource.SinaSimple
	err = json.Unmarshal([]byte(dataStrDecode), &detail)
	if err != nil {
		seelog.Errorf("index解析错误，err = %+v\n", err)
		return
	}
	for _, v := range detail {
		fmt.Printf("code = %+v\n", v.Code)
		fmt.Printf("name = %+v\n", v.Name)
	}
	return
}

// 检查分类索引
func checkIndex(body []byte) (ret map[string]string) {
	dataStr := string(body)
	strDecode, err := iconv.ConvertString(dataStr, "GB2312", "utf-8")
	if err != nil {
		seelog.Errorf("err = %+v\n", err)
		return
	}
	data := strings.Split(strDecode, "=")
	if len(data) >= 2 {
		var dataDetail map[string]string
		err := json.Unmarshal([]byte(data[1]), &dataDetail)
		if err != nil {
			seelog.Errorf("json解析错误，err = %+v\n", err)
			return
		}
		for tag, v := range dataDetail {
			vDetail := strings.Split(string(v), ",")
			ret[tag] = vDetail[1]
		}
	}
	return
}

// CheckIndustry 检测产业分类，并储存
func (api ClassifyAPI) CheckIndustry() (err error) {
	host := library.SchemeHttp + library.URLSinaIndustryIndex
	body, err := api.Config.doGet(host)
	if err != nil {
		seelog.Errorf("err = %+v\n", err)
		return err
	}
	tagMap := checkIndex(body)
	for tag, classifyName := range tagMap {
		checkDetailByTag(tag, classifyName, api.Config)
	}
	return
}

// CheckConcept 检测概念分类，并储存
func (api ClassifyAPI) CheckConcept() (err error) {
	host := library.SchemeHttp + library.URLSinaConceptIndex
	body, err := api.Config.doGet(host)
	if err != nil {
		seelog.Errorf("err = %+v\n", err)
		return err
	}
	tagMap := checkIndex(body)
	for tag, classifyName := range tagMap {
		checkDetailByTag(tag, classifyName, api.Config)
	}
	return
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
		api.Config.UriPath = library.URLClassifyArea
		api.Config.DataStruct = &resource.Area{}
	// 获取概念(获取数据频率太快，会导致被墙)
	case Concept:
		api.Config.UriPath = library.URLClassifyConcept
		api.Config.DataStruct = &resource.Concept{}
	// 获取行业(获取数据频率太快，会导致被墙)
	case Industry:
		api.Config.UriPath = library.URLClassifyIndustry
		api.Config.DataStruct = &resource.Industry{}
	// 获取中小板分类
	case SME:
		api.Config.UriPath = library.URLClassifySME
		api.Config.DataStruct = &resource.SME{}
	// 获取创业板分类
	case GEM:
		api.Config.UriPath = library.URLClassifyGEM
		api.Config.DataStruct = &resource.GEM{}
	// 获取风险警示板
	case ST:
		api.Config.UriPath = library.URLClassifyST
		api.Config.DataStruct = &resource.ST{}
	// 获取沪深300
	case HS300:
		api.Config.UriPath = library.URLClassifyHS300
		api.Config.DataStruct = &resource.HS300{}
	// 获取上证50
	case SZ50:
		api.Config.UriPath = library.URLClassifySZ50
		api.Config.DataStruct = &resource.SZ50{}
	// 获取中证500
	case ZZ500:
		api.Config.UriPath = library.URLClassifyZZ500
		api.Config.DataStruct = &resource.ZZ500{}
	// 获取暂停上市
	case Suspended:
		api.Config.UriPath = library.URLClassifySuspended
		api.Config.DataStruct = &resource.Suspended{}
	// 获取终止上市
	case Terminated:
		api.Config.UriPath = library.URLClassifyTerminated
		api.Config.DataStruct = &resource.Terminated{}
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

func (api ClassifyAPI) GateJob(c *gin.Context) {

}
