package main

import (
	ts "assets/web/backend/tushare"

	seelog "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置seelog
	defer seelog.Flush()
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			seelog.Errorf("err = %+v\n", err)
		}
	}()
	logger, err := seelog.LoggerFromConfigAsFile("./conf/seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/deal/:gateway", ts.NewDealAPI().GateWay)
		v1.POST("/bank/:gateway", ts.NewBankAPI().GateWay)
		v1.POST("/bill/:gateway", ts.NewBillBoardAPI().GateWay)
		v1.POST("/classify/:gateway", ts.NewClassifyAPI().GateWay)
		v1.POST("/company/:gateway", ts.NewCompanyAPI().GateWay)
		v1.POST("/economic/:gateway", ts.NewEconomicAPI().GateWay)
		v1.POST("/invest/:gateway", ts.NewInvestAPI().GateWay)
		v1.POST("/news/:gateway", ts.NewNewsAPI().GateWay)

		v1.GET("/classify/job/:gatejob", ts.NewClassifyAPI().GateJob)
	}
	r.Run(":8000")
}
