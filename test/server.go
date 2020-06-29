package main

import (
	"github.com/genji1037/gin-cache-middle"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	engine := gin.New()
	failedFn := func(c *gin.Context) {
		c.JSON(500, "internal server error")
	}
	engine.Use(gin_cache_middle.GetCacheMiddle(&gin_cache_middle.MockCache{}, failedFn, 5*time.Second))
	engine.GET("/hello", func(c *gin.Context) {
		time.Sleep(500 * time.Millisecond)
		c.JSON(200, map[string]interface{}{
			"rand": rand.Int(),
		})
	})
	engine.Run(":8123")
}
