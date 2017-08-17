package library

import (
	"errors"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
)

func CheckUrlParams(params url.Values) error {
	var err error
	err = nil
	for k, v := range params {
		// 校验字符串
		if v[0] == "" {
			err = errors.New(k + "的参数为空")
			break
		}
		// 校验数组
	}
	return err
}

func CallFuncByName(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("函数参数的数量错误")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func OutputErr(c *gin.Context, err error, staCode int) {
	c.JSON(staCode, gin.H{
		"data": struct{}{},
		"msg":  err.Error(),
	})
	return
}
