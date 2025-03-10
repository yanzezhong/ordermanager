package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mongo MongoConfig // mongo 相关配置
	Auth  struct {    // Key and expiration time configuration required for JWT authentication
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}
}

type MongoConfig struct {
	Url string
	DB  string `json:",default=jarvis"`
}
