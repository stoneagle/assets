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
	r.GET("/bank/getShibor/:year", ts.NewBankAPI().GetShibor)
	r.POST("/deal/getKData/:code", ts.NewDealAPI().GetKData)
	r.Run(":8000")
}
