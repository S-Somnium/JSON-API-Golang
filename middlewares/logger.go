package middlewares

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		raw := c.Request.URL.RawQuery
		c.Next()
		if raw != "" {
			path = path + "?" + raw
		}
		log.WithFields(log.Fields{
			"StatusCode": c.Writer.Status(),
			"Path":       path,
			"Method":     c.Request.Method,
			"Duration":   fmt.Sprint(time.Since(start).Milliseconds()) + "ms",
			"ClientIP":   c.ClientIP(),
		}).Info("Incoming request")
	}
}
