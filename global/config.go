package global

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	// 日志路径
	Log struct {
		LogDir string `yaml:"logDir"`
	} `yaml:"log"`
	// 数据库连接
	DBConn struct {
		Address      string `yaml:"address"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DBName       string `yaml:"dbName"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	} `yaml:"dbConn"`
	RedisConn struct {
		Address     string `yaml:"address"`
		IdleTimeout int    `yaml:"idleTimeout"`
		MaxIdle     int    `yaml:"maxIdle"`
		MaxActive   int    `yaml:"maxActive"`
		Password    string `yaml:"password"`
	} `yaml:"redisConn"`
}

var Config *config

func init() {
	Config = &config{}
	configPath := flag.String("c", "./config.yml", "config path")
	flag.Parse()
	content, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, Config)
	if err != nil {
		panic(err)
	}
}
