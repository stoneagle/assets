package tushare

import (
	"assets/web/backend/library"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/cihub/seelog"
)

type ResponseData struct {
	Errno  string                   `json:"errno"`
	Errmsg string                   `json:"errmsg"`
	Data   []map[string]interface{} `json:"data"`
}

type ResponseFailData struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type Configuration struct {
	BasePath   string `json:"basePath,omitempty"`
	UriPath    string `json:"uriPath,omitempty"`
	DataStruct interface{}
	Params     url.Values
}

func NewConfig() *Configuration {
	return &Configuration{
		BasePath: "http://django:8000",
		UriPath:  "",
	}
}

func (c Configuration) doHttp() (ret *ResponseData, err error) {
	err = library.CheckUrlParams(c.Params)
	if err != nil {
		seelog.Infof("参数校验结果，err:%s", err)
		return
	}

	client := &http.Client{}

	host := c.BasePath + c.UriPath
	req, err := http.NewRequest("POST", host, strings.NewReader(c.Params.Encode()))
	if err != nil {
		seelog.Infof("http请求准备失败，url:%s", host)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		seelog.Infof("req执行失败")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		seelog.Infof("ioutil读取失败，url:%s", host)
		return
	}

	if resp.StatusCode != http.StatusOK {
		failRet := &ResponseFailData{}
		err = json.Unmarshal(body, &failRet)
		if err == nil {
			err = errors.New("tushare资源错误[" + failRet.Errno + "]:" + failRet.Errmsg)
		}
	} else {
		err = json.Unmarshal(body, &ret)
	}
	return
}
