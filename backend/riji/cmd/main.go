package main

import (
	"flag"
	"riji/api"
	"riji/service"

	"github.com/gin-gonic/gin"
	"github.com/qq1141000259/public_tools/zlog"
)

var (
	addr    string = ":8099"
	log_dir string = ""
)

func init() {
	flag.StringVar(&addr, "addr", addr, "server listen address")
}

func main() {
	flag.Parse()

	// 初始化zlog
	logCfg := zlog.NewZapConfig()
	logCfg.LoggerType = "json"
	logCfg.Alarm = false
	logCfg.ServerName = "riji-go"
	logCfg.Build()

	// 初始化Server
	s, err := service.NewServer()
	if err != nil {
		zlog.Fatal(err.Error())
	}

	r := gin.New()
	api.Install(r, s)
	r.Run(addr)
}
