package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	logger "github.com/isayme/go-logger"
)

var config Config

// Config service config
type Config struct {
	HTTP struct {
		Port int `json:"port"`
	} `json:"http"`

	Logger struct {
		Level string `json:"level"`
	} `json:"logger"`

	Mongo Mongo `json:"mongo"`
	JWT   JWT   `json:"jwt"`
}

// Mongo mongo config
type Mongo struct {
	Addrs          []string `json:"addrs"`
	Timeout        int      `json:"timeout"` // in seconds
	Database       string   `json:"database"`
	ReplicaSetName string   `json:"replicaSet"`
	Source         string   `json:"source"`
	PoolLimit      int      `json:"poolLimit"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
}

// JWT jwt config
type JWT struct {
	Method string   `json:"method"`
	Keys   []string `json:"keys"`
	Expire Duration `json:"expire"`
}

var once sync.Once

// Get get config
func Get() *Config {
	once.Do(func() {
		path := os.Getenv("CONF_FILE_PATH")
		logger.Debugf("config file path: %s", path)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(data, &config); err != nil {
			panic(err)
		}
	})

	return &config
}
