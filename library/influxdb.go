package library

import (
	"fmt"
	"log"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	// BitmexDB 数据库名
	BitmexDB = "bitmex"
	// TushareDB 数据库名
	TushareDB = "tushare"
	username  = "wuzhongyang"
	password  = "a1b2c3d4E"
)

// InfluxdbAPI 对象
type InfluxdbAPI struct {
	Client client.Client
	DBName string
}

// NewInfluxdbAPI 新建InfluxdbAPI对象
func NewInfluxdbAPI(dbName string) *InfluxdbAPI {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://influxdb.localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	influxdbAPI := &InfluxdbAPI{Client: c, DBName: dbName}
	return influxdbAPI
}

// WritePoints 插入数据
func (api InfluxdbAPI) WritePoints(points []client.Point) error {
	batchPoint, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  api.DBName,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, point := range points {
		batchPoint.AddPoint(&point)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = api.Client.Write(batchPoint)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (api InfluxdbAPI) queryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: api.DBName,
	}
	if response, err := api.Client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err

	}
	return res, nil
}

// CreateDB 创建数据库
func (api InfluxdbAPI) CreateDB() error {
	_, err := api.queryDB(fmt.Sprintf("CREATE DATABASE %s", api.DBName))
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// GetCount 获取计数
func (api InfluxdbAPI) GetCount(field string, tableName string) (interface{}, error) {
	query := fmt.Sprintf("SELECT count(%s) FROM %s", field, tableName)
	res, err := api.queryDB(query)
	if err != nil {
		log.Fatal(err)
	}
	count := res[0].Series[0].Values[0][1]
	return count, err
}

// GetList 获取列表
func (api InfluxdbAPI) GetList(tableName string, limit int) ([][]interface{}, error) {
	var query string
	if limit <= 0 {
		query = fmt.Sprintf("SELECT * FROM %s", tableName)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s LIMIT %d", tableName, limit)
	}
	res, err := api.queryDB(query)
	if err != nil {
		log.Fatal(err)
	}
	return res[0].Series[0].Values, err
}
