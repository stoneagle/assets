package bitmex

import (
	"assets/web/backend/bitmex/model"
	"assets/web/backend/library"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb/client/v2"
)

// HistoryAPI 对象
type HistoryAPI struct {
	influxdb library.InfluxdbAPI
}

// NewHistoryAPI 创建HistoryAPI对象
func NewHistoryAPI() *HistoryAPI {
	influxdb := library.NewInfluxdbAPI(library.BitmexDB)
	return &HistoryAPI{influxdb: *influxdb}
}

// GetData 获取分钟级别数据
// func GetData(string resolution, string symbol, int64 from, int64 to) {
func (a HistoryAPI) getData() (*model.HistoryData, error) {
	url := "https://www.bitmex.com/api/udf/history?symbol=XRPM17&resolution=5&from=1497359200&to=1497359852"
	historyData := new(model.HistoryData)
	var rawData []byte

	response, err := http.Get(url)
	if err != nil {
		return historyData, err
	}

	rawData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return historyData, err
	}

	err = json.Unmarshal(rawData, historyData)
	return historyData, err
}

// SpiderHistory 获取历史数据
func (a HistoryAPI) SpiderHistory(c *gin.Context) {
	var err error
	historyData := new(model.HistoryData)
	historyData, err = a.getData()
	if err != nil {
		log.Printf("response--------------%v", historyData)
		c.JSON(200, gin.H{
			"code": "0001",
			"data": "",
			"msg":  err,
		})
	} else {
		var points []client.Point
		for i := 0; i < len(historyData.Time); i++ {
			fmt.Println(historyData.Time[i])
			tags := map[string]string{
				"symbol": "XRPM17",
			}
			fields := map[string]interface{}{
				"time":   historyData.Time[i],
				"close":  historyData.Close[i],
				"open":   historyData.Open[i],
				"low":    historyData.Low[i],
				"high":   historyData.High[i],
				"volumn": historyData.Volumn[i],
			}
			point, err := client.NewPoint(model.HistoryDayTable, tags, fields, time.Now())
			points = append(points, *point)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = a.influxdb.WritePoints(points)
		c.JSON(200, gin.H{
			"code": "0000",
			"data": historyData,
			"msg":  "success",
		})
	}
}

// GetHistoryList 获取历史数据列表
func (a HistoryAPI) GetHistoryList(c *gin.Context) {
	var result [][]interface{}
	var err error
	result, err = a.influxdb.GetList(model.HistoryDayTable, 0)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"code": "0000",
		"data": result,
		"msg":  "success",
	})
}
