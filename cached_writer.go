package gin_cache_middle

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type cachedWriter struct {
	gin.ResponseWriter
	cache  Cache
	expire time.Duration
	key    string
	Body   []byte
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		w.Body = data
		err = w.cache.Set(w.key, w.Body, w.expire)
		if err != nil {
			log.Printf("cachedWriter cache set failed: %s", err.Error())
		}
	}
	return ret, err
}
