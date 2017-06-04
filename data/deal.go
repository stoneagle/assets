package data

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// 历史行情
func getHistoryData(c *gin.Context) {

}

// 复权
func getBasicData(c *gin.Context) {

}

// 实时
func getTodayData(c *gin.Context) {

}

// 历史分笔
func getHistoryTick(c *gin.Context) {

}

// 当日分笔
func getTodayTick(c *gin.Context) {

}

// 大盘指数
func getIndex(c *gin.Context) {

}

// 大单交易数据
func getBigDeal(c *gin.Context) {

}
