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
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ResponseData struct {
	Errno  string                   `json:"errno"`
	Errmsg string                   `json:"errmsg"`
	Data   []map[string]interface{} `json:"data,omitempty"`
}

type ResponseCheck struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type Configuration struct {
	BasePath   string `json:"basePath,omitempty"`
	UriPath    string `json:"uriPath,omitempty"`
	DataStruct interface{}
	Params     url.Values
	SqlDB      *sqlx.DB
}

func NewConfig() *Configuration {
	db, err := sqlx.Connect("mysql", "root:z20138502@wzy-360@tcp(123.207.151.42:3306)/asset")
	if err != nil {
		seelog.Errorf("db初始化失败 err = %+v\n", err)
	}
	return &Configuration{
		BasePath: "http://django:8000",
		UriPath:  "",
		SqlDB:    db,
	}
}

func (c Configuration) doPost() (ret *ResponseData, err error) {
	err = library.CheckUrlParams(c.Params)
	if err != nil {
		seelog.Infof("参数校验结果，err:%s", err)
		return
	}

	client := &http.Client{}
	host := c.BasePath + c.UriPath

	seelog.Debug("host:%s", host)
	seelog.Debug("Params:%v", c.Params)

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

	check := &ResponseCheck{}
	err = json.Unmarshal(body, &check)
	if err != nil {
		seelog.Infof("tushare的check读取失败，ret:%s", string(body))
		return
	}

	if check.Errno == "0000" {
		err = json.Unmarshal(body, &ret)
	} else {
		err = errors.New("tushare资源错误[" + check.Errno + "]:" + check.Errmsg)
	}
	return
}

func (c Configuration) doGet(host string) (body []byte, err error) {
	resp, err := http.Get(host)
	if err != nil {
		seelog.Infof("req执行失败")
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		seelog.Infof("ioutil读取失败，url:%s", host)
		return
	}
	return
}
