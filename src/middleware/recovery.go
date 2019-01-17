package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-user/src/httperror"
	"github.com/pkg/errors"
)

// Recovery recover from panic
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			switch v := err.(type) {
			case *httperror.HTTPError:
				c.AbortWithStatusJSON(v.Status, gin.H{
					"name":    v.Name,
					"message": v.Error(),
				})
			default:
				// httprequest, _  := httputil.DumpRequest(c.Request, false)
				// logger.Warnf("[Recovery] %s panic recovered:\n%s\n%s", timeFormat(time.Now()), string(httprequest), err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"name":    "Error",
					"message": fmt.Sprintf("%v", err),
					"stack":   fmt.Sprintf("%+v", errors.WithStack(err.(error))),
				})
			}
		}
	}()
	c.Next()
}
