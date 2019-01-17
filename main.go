package main

import (
	"flag"
	"fmt"
	"os"

	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-user/src"
	"github.com/isayme/go-user/src/conf"
	"github.com/isayme/go-user/src/di"
	"github.com/isayme/go-user/src/util"
)

var showVersion = flag.Bool("v", false, "show version")
var showHelp = flag.Bool("h", false, "show help")

func main() {
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("name: %s\nversion: %s\nbuildTime: %s\ngitRevision: %s\n", util.Name, util.Version, util.BuildTime, util.GitRevision)
		os.Exit(0)
	}

	cfg := conf.Get()
	if cfg.Logger.Level != "" {
		if err := logger.SetLevel(cfg.Logger.Level); err != nil {
			panic(err)
		}
	}

	err := di.Invoke(func(app *src.App) error {
		logger.Panic(app.Run())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
