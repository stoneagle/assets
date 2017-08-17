package main

import (
	ts "assets/web/backend/tushare"

	seelog "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置seelog
	defer seelog.Flush()
	logger, err := seelog.LoggerFromConfigAsFile("./conf/seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)

	r := gin.Default()
	// r.GET("/announcement", bitmex.NewAnnouncementAPI().GetList)
	v1 := r.Group("/v1")
	{
		// v1Deal.POST("/getKData", ts.NewDealAPI().GetKData)
		v1.POST("/deal/:gateway", ts.NewDealAPI().GateWay)
		v1.POST("/bank/:gateway", ts.NewBankAPI().GateWay)
	}
	r.Run(":8000")
}
