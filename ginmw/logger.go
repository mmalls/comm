package ginmw

import (
	"math"
	"os"
	"time"

	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/xtfly/log4g"
)

// Logger is the logrus logger handler
func Logger(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()

		if logger.DebugEnabled() {
			bs, _ := httputil.DumpRequest(c.Request, true)
			logger.Debug(string(bs))
		}

		c.Next()

		if logger.DebugEnabled() {
			stop := time.Since(start)
			latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000.0))
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			referer := c.Request.Referer()
			hostname, err := os.Hostname()
			if err != nil {
				hostname = "unknown"
			}
			dataLength := c.Writer.Size()
			if dataLength < 0 {
				dataLength = 0
			}

			if len(c.Errors) > 0 {
				logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				logger.Debugf("[%s][%s][%s %s][%d][%d][%s][%s][%dms]",
					clientIP, hostname, c.Request.Method, path, statusCode, dataLength, referer, latency)
			}
		}
	}
}
