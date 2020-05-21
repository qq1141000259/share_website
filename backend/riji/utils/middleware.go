package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qq1141000259/public_tools/zlog"
)

// 中间键--错误处理+api日志
func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		defer func() {
			// 记录接口耗时
			duration := time.Now().Sub(now).Seconds()
			if err := recover(); err != nil {
				var (
					errMsg string
					ok     bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error, " + errMsg,
					})
					zlog.Errorf("URL: %s duration: %fs msg: %s", c.Request.URL, duration, "system error, "+errMsg)
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error",
					})
					zlog.Errorf("URL: %s duration: %fs msg: %s %+v", c.Request.URL, duration, "system error", err)
					return
				}
			}
		}()
		// 解析请求的json参数
		data, _ := ioutil.ReadAll(c.Request.Body)
		// 因为request.body是readcloser 需要重新填充Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		c.Next()
		zlog.Infof("URL: %s Req: %s Rsp: %s", c.Request.URL, string(data))
	}
}

type handleFunc func(c *gin.Context)

// 登录中间件
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过对 "api/user 类接口的检查"
		if !strings.Contains(c.Request.URL.String(), "api/user") {
			session := sessions.Default(c)
			v := session.Get("user_id")
			if v == nil || v == 0 {
				ValidateError(c, 99999, "登录失效，重新登录")
				c.Abort()
				return
			} else {
				c.Set("user_id", v.(uint32))
				c.Set("user_name", session.Get("user_name").(string))
			}
			c.Next()
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.String(200, "ok")
			return
		}
		c.Next()
	}
}
