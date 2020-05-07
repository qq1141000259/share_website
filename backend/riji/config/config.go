package config

import (
	"time"
)

// Config 模块配置 需要自行按需修改
type Config struct {
	*DB            `mapstructure:"db"`
	*CosConfig     `mapstructure:"cos"`
}

// DB mysql config.
type DB struct {
	DSN         string        `mapstructure:"dsn"`          // data source name.
	Active      int           `mapstructure:"active"`       // pool
	Idle        int           `mapstructure:"idle"`         // pool
	IdleTimeout time.Duration `mapstructure:"idle_timeout"` // connect max life time.
}

//
type CosConfig struct {
	Bucket string `mapstructure:"bucket"`
	Si string `mapstructure:"si"`
	Sk string `mapstructure:"sk"`
}

