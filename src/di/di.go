package di

import (
	"github.com/isayme/go-user/src"
	"github.com/isayme/go-user/src/conf"
	"github.com/isayme/go-user/src/dao"
	"github.com/isayme/go-user/src/mongo"
	"github.com/isayme/go-user/src/router"
	"go.uber.org/dig"
)

var defaultContainer = dig.New()

// Provide ...
var Provide = defaultContainer.Provide

// Invoke ...
var Invoke = defaultContainer.Invoke

func init() {
	Provide(func() *conf.Config {
		return conf.Get()
	})
	Provide(mongo.NewSession)
	Provide(dao.NewUser)
	Provide(router.NewUser)
	Provide(src.NewApp)
}
