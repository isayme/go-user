package src

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/isayme/go-user/src/conf"
	"github.com/isayme/go-user/src/dao"
	"github.com/isayme/go-user/src/middleware"
)

// App ...
type App struct {
	config *conf.Config
	r      *gin.Engine
	user   *dao.User
}

// NewApp create app
func NewApp(config *conf.Config, user *dao.User) *App {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)

	app := &App{
		config: config,
		r:      r,
		user:   user,
	}
	registerRouter(app)

	return app
}

// Run serve
func (app *App) Run() error {
	return app.r.Run(fmt.Sprintf(":%d", app.config.HTTP.Port))
}

// func (app *App) Close(ctx ...context.Context) error {
// 	return app.r.
// }
