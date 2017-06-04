package main

import (
	"fmt"

	"assets/data"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	fmt.Println("test")
	r := gin.Default()
	r.GET("/ping", data.Ping)
	r.Run(":6565")
}
