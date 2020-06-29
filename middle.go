package gincachemiddle

import (
	"time"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
	"log"
)

var g singleflight.Group

func GetCacheMiddle(cache Cache, failedRespond func(c *gin.Context), expire time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		body, err := g.Do(url, func() (i interface{}, e error) {
			value, ok := cache.Get(url)
			if ok {
				c.Writer.WriteHeader(200)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.Write(value)
				return nil, nil
			}
			cacheWriter := &cachedWriter{
				ResponseWriter: c.Writer,
				cache:          cache,
				expire:         expire,
				key:            url,
			}
			c.Writer = cacheWriter
			c.Next()

			return cacheWriter.Body, nil
		})
		if err != nil {
			log.Printf("cache middle failed: %s\n", err.Error())
			failedRespond(c)
		} else if !c.Writer.Written() {
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.Write(body.([]byte))
		}
		c.Abort()
		return
	}
}
