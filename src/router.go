package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-user/src/middleware"
	"github.com/isayme/go-user/src/router"
	"github.com/isayme/go-user/src/util"
)

func registerRouter(app *App) {
	app.r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    util.Name,
			"version": util.Version,
		})
	})

	v1 := app.r.Group("/v1")
	userRouter := router.NewUser(app.user)
	v1.POST("/users/signup", userRouter.Signup)
	v1.POST("/users/login", userRouter.Login)
	v1.Use(middleware.AuthorizeRequired())
	v1.GET("/users/me", userRouter.Me)
}
