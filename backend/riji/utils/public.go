package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": 404,
		"msg":  "找不到该路由",
	})
}

func PaginationHandler(p string, limit string) (pNum, limitNum int) {
	defaultP := 1
	defaultL := 9

	var err error
	if p == "" {
		pNum = defaultP
	} else {
		pNum, err = strconv.Atoi(p)
		if err != nil || pNum <= 0 {
			pNum = defaultP
		}
	}

	if limit == "" {
		limitNum = defaultL
	} else {
		limitNum, err = strconv.Atoi(limit)
		if err != nil || limitNum <= 0 {
			limitNum = defaultL
		}
	}
	return
}
