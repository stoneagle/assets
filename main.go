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
	r.Run(":8000")
}
