package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// httprequest, _  := httputil.DumpRequest(c.Request, false)
			// logger.Warnf("[Recovery] %s panic recovered:\n%s\n%s", timeFormat(time.Now()), string(httprequest), err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
				"stack": fmt.Sprintf("%+v", errors.WithStack(err.(error))),
			})
		}
	}()
	c.Next()
}
