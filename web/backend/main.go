package main

import (
	"assets/bitmex"
	"assets/tushare"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", tushare.Ping)
	r.GET("/announcement", bitmex.NewAnnouncementAPI().GetList)
	r.GET("/apikey", bitmex.NewAPIKeyAPI().GetList)
	r.GET("/spider", bitmex.NewHistoryAPI().SpiderHistory)
	r.GET("/history", bitmex.NewHistoryAPI().GetHistoryList)
	r.Run(":8000")
}
