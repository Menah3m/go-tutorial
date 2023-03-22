package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		t := time.Now()
		c.Next() // 先继续处理其他请求
		logger.Info("incoming request:",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status code", c.Writer.Status()),
			zap.Duration("latency", time.Now().Sub(t)),
		)

	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "ping",
		}
		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}

		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run(":8080")

}
