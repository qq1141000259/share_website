package service

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"riji/model"
	. "riji/utils"

	"github.com/gin-gonic/gin"
	"github.com/qq1141000259/public_tools/httpclient"
)

// 敲出一个hello world
func (s *RijiServer) HelloWorld(c *gin.Context) {
	var res []*model.BackImgCos
	var total int
	if err := s.dao.GetTablePagination(&res, &total, 1, 20, nil); err != nil {
		ValidateError(c, DBError, err.Error())
	}
	var insertRow model.BackImgCos
	insertRow.AddBy = 1
	if err := s.dao.CreateBackImageCos(&insertRow); err != nil {
		ValidateError(c, DBError, err.Error())
	}
	Response(c, res)
}

// 更新
func (s *RijiServer) UpdateBackImg(c *gin.Context) {
	client := httpclient.NewClient()
	rsp, err := client.GET(BingBackImgUrl, nil)
	if err != nil {
		ValidateError(c, OtherHttpError, "请求bing网站失败")
		return
	}
	defer rsp.Body.Close()
	// 解析返回
	body, errRead := ioutil.ReadAll(rsp.Body)
	if errRead != nil {
		ValidateError(c, CodeError, "请求bing网站失败")
		return
	}
	bodyStr := string(body)
	rex, err := regexp.Compile(".*data-ultra-definition-src=\"(.*)\" data-explicit-bing-load.*")
	if err != nil {
		ValidateError(c, CodeError, "请求bing网站失败")
		return
	}
	target := rex.FindStringSubmatch(bodyStr)[1]
	fmt.Println(target)
	Response(c, target)
}

// 获取最新的背景图片地址
func (s *RijiServer) GetBackImgUrl(c *gin.Context) {
	cos := NewCos(s.conf.Bucket, s.conf.Si, s.conf.Sk)
	var backImg []*model.BackImgCos
	if err := s.dao.GetTableOrderByCreateTime(&backImg); err != nil {
		ValidateError(c, DBError, fmt.Sprintf("查询失败 %s", err.Error()))
		return
	}
	var urls []string
	for _, obj := range backImg {
		url, err := cos.GetDownloadUrl(obj.ID + "." + obj.Suffix)
		if err != nil {
			ValidateError(c, CosError, fmt.Sprintf("cos调用失败 %s", err.Error()))
			return
		}
		urls = append(urls, url)
	}
	Response(c, urls)
}
