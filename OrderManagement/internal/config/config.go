package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mongo          MongoConfig     // mongo 相关配置
}


type MongoConfig struct {
	Url string
	DB  string `json:",default=jarvis"`
}