package router

import (
	"github.com/gin-gonic/gin"
	"time"
)

func PubInitRouter(engine *gin.Engine) {
	r := engine.Group("/pub")
	r.GET("/ping", PubGetServerTime)
	r.GET("/panic", func(ctx *gin.Context) {
		panic("/pub/panic")
	})
}

func PubGetServerTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"now": time.Now(),
		},
	})
}
