package service

import (
	"riji/model"
	. "riji/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageCreateReq struct {
	Content string `json:"content"`
}

type CosCallbackReq struct {
	Key    string `json:"key"`
	Suffix string `json:"suffix"`
	FType  string `json:"type"`
	Title  string `json: "title"`
}

// 删除图片
func (s *RijiServer) PicDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	var table model.UserUploadFile
	if err := s.dao.GetTableById(&table, id); err != nil {
		ValidateError(c, RecordNotFound, err.Error())
	}
	if err := s.dao.UpdateTable(&table, map[string]bool{"c_is_delete": true}); err != nil {
		ValidateError(c, DBError, err.Error())
	}
	Response(c, "删除成功")
}

// 图片列表
func (s *RijiServer) PicList(c *gin.Context) {
	p, limit := PaginationHandler(c.DefaultQuery("p", "1"), c.DefaultQuery("limit", "9"))
	var tables []*model.UserUploadFile
	var total int
	conditions := map[string]interface{}{
		"c_upload_type": model.PhotoType,
	}

	if err := s.dao.GetTablePagination(&tables, &total, p, limit, conditions); err != nil {
		ValidateError(c, RecordNotFound, err.Error())
		return
	}

	// 获取图片地址
	cos := NewCos(s.conf.Bucket, s.conf.Si, s.conf.Sk)
	var urls []map[string]string
	var title string
	for _, obj := range tables {
		url, err := cos.GetDownloadUrl(obj.ID + "." + obj.FileType)
		if err != nil {
			ValidateError(c, CosError, err.Error())
			return
		}
		if obj.Title == "" {
			title = "未命名"
		} else {
			title = obj.Title
		}
		urls = append(urls, map[string]string{
			"url":   url,
			"title": title,
		})
	}

	Response(c, map[string]interface{}{"rows": urls, "total": total})
}

// 视频文件列表
func (s *RijiServer) VideoList(c *gin.Context) {
	p, limit := PaginationHandler(c.Query("p"), c.Query("limit"))
	var tables []*model.UserUploadFile
	var total int
	conditions := map[string]interface{}{
		"c_upload_type": model.VideoType,
	}

	if err := s.dao.GetTablePagination(&tables, &total, p, limit, conditions); err != nil {
		ValidateError(c, RecordNotFound, err.Error())
		return
	}

	// 获取图片地址
	cos := NewCos(s.conf.Bucket, s.conf.Si, s.conf.Sk)
	var urls []map[string]string
	var title string
	for _, obj := range tables {
		url, err := cos.GetDownloadUrl(obj.ID + "." + obj.FileType)
		if err != nil {
			ValidateError(c, CosError, err.Error())
			return
		}
		if obj.Title == "" {
			title = "未命名"
		} else {
			title = obj.Title
		}
		urls = append(urls, map[string]string{
			"url":   url,
			"title": title,
		})
	}

	Response(c, map[string]interface{}{"rows": urls, "total": total})
}

// 获取留言板列表
func (s *RijiServer) MessageList(c *gin.Context) {
	p, limit := PaginationHandler(c.Query("p"), c.Query("limit"))
	var tables []*model.MessageBoard
	var total int

	if err := s.dao.GetTablePagination(&tables, &total, p, limit, nil); err != nil {
		ValidateError(c, RecordNotFound, err.Error())
		return
	}
	var data []interface{}
	for _, obj := range tables {
		data = append(data, map[string]string{
			"id":      obj.ID,
			"date":    obj.CreatedAt.Format("2006-01-02 15:04:05"),
			"content": obj.Content,
			"user":    s.dao.GetUserName(obj.AddBy),
		})
	}
	Response(c, map[string]interface{}{
		"rows":  data,
		"total": total,
	})
}

// 新增留言
func (s *RijiServer) MessageCreate(c *gin.Context) {
	var table model.MessageBoard
	var req MessageCreateReq
	if err := c.BindJSON(&req); err != nil {
		ValidateError(c, ParamError, err.Error())
		return
	}
	if req.Content == "" {
		ValidateError(c, ParamError, "内容不能为空")
		return
	}
	table.Content = req.Content
	table.AddBy = c.Keys["user_id"].(uint32)

	if err := s.dao.CreateTable(&table); err != nil {
		ValidateError(c, DBError, err.Error())
	}
	Response(c, "新增成功")
}

// 删除留言
func (s *RijiServer) MessageDelete(c *gin.Context) {
	msgId := c.Params.ByName("id")
	if msgId == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	var table model.MessageBoard
	if err := s.dao.GetTableById(&table, msgId); err != nil {
		ValidateError(c, DBError, err.Error())
		return
	}
	attrs := map[string]interface{}{
		"c_update_by": c.Keys["user_id"].(uint32),
		"c_is_delete": true,
	}
	if err := s.dao.UpdateTable(&table, attrs); err != nil {
		ValidateError(c, DBError, err.Error())
	}
	Response(c, "删除成功")
}

// 生成cos上传凭证
func (s *RijiServer) GetCredentials(c *gin.Context) {
	cos := NewCos(s.conf.Bucket, s.conf.Si, s.conf.Sk)
	cr, err := cos.CreateCredentials()
	if err != nil {
		ValidateError(c, CosError, err.Error())
		return
	}
	now := time.Now().Unix()
	ext := time.Now().Add(time.Second * 3600).Unix()
	data := map[string]interface{}{
		"TmpSecretId":  cr.TmpSecretID,
		"TmpSecretKey": cr.TmpSecretKey,
		"Token":        cr.SessionToken,
		"StartTime":    now,
		"ExpiredTime":  ext,
	}
	Response(c, data)
}

// cos文件上传成功，回调
func (s *RijiServer) CosCallback(c *gin.Context) {
	var dat CosCallbackReq
	if err := c.BindJSON(&dat); err != nil {
		ValidateError(c, ParamError, err.Error())
		return
	}
	if dat.Key == "" || dat.FType == "" || dat.Suffix == "" {
		ValidateError(c, ParamError, "参数不能为空")
		return
	}
	var table model.UserUploadFile
	table.ID = dat.Key
	table.FileType = dat.Suffix
	table.Title = dat.Title
	table.AddBy = c.Keys["user_id"].(uint32)
	if dat.FType == "video" {
		table.UploadType = model.VideoType
	} else {
		table.UploadType = model.PhotoType
	}
	if err := s.dao.CreateTable(&table); err != nil {
		ValidateError(c, DBError, err.Error())
		return
	}
	Response(c, "上传成功")
	return
}
