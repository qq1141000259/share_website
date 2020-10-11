package service

import (
	"github.com/gin-gonic/gin"
	"riji/model"
	. "riji/utils"
	"strings"
)

type BlogPostReq struct {
	Title   string   `json:"title"`
	Tag     []string `json:"tag"`
	Desc    string   `json:"desc"`
	Content string   `json:"content"`
}

func (s *RijiServer) BlogGet(c *gin.Context) {
	p, limit := PaginationHandler(c.Query("p"), c.Query("limit"))
	var tables []*model.Blog
	var count int
	// 按条件查询
	conditions := make(map[string]interface{})
	//if c.Query("user") != "" {
	//	conditions["c_word"] = c.Query("word")
	//}
	//if c.Query("annotation") != "" {
	//	conditions["c_annotation"] = "%" + c.Query("annotation") + "%"
	//}
	//if c.Query("remark") != "" {
	//	conditions["c_remark"] = "%" + c.Query("remark") + "%"
	//}
	if err := s.dao.GetTablePagination(&tables, &count, p, limit, conditions); err != nil {
		ValidateError(c, DBError, "查询博客记录失败")
		return
	}
	var data []map[string]interface{}
	for _, obj := range tables {
		data = append(data, map[string]interface{}{
			"id":    obj.ID,
			"title": obj.Title,
			"tag":   strings.Split(obj.Tag, ","),
			"desc":  obj.Desc,
			"date":  obj.CreatedAt.Format("2006-01-02"),
			"user":  s.dao.GetUserName(obj.AddBy),
		})
	}
	Response(c, map[string]interface{}{
		"rows":  data,
		"total": count,
	})
}

func (s *RijiServer) BlogPost(c *gin.Context) {
	var req BlogPostReq
	if err := c.BindJSON(&req); err != nil {
		ValidateError(c, ParamError, "请求参数格式错误")
		return
	}
	if req.Title == "" || req.Content == "" {
		ValidateError(c, ParamError, "缺少参数标题或正文")
		return
	}
	table := model.Blog{
		Title:   req.Title,
		Tag:     strings.Join(req.Tag, ""),
		Desc:    req.Desc,
		Content: req.Content,
	}
	table.AddBy = c.Keys["user_id"].(uint32)
	if err := s.dao.CreateTable(&table); err != nil {
		ValidateError(c, DBError, "插入数据失败")
		return
	}
	Response(c, "添加文章成功")
}

func (s *RijiServer) BlogPut(c *gin.Context) {
	var req BlogPostReq
	blogId := c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		ValidateError(c, ParamError, "请求参数格式错误")
		return
	}
	if blogId == "" || req.Title == "" || req.Content == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	table := model.Blog{}
	if err := s.dao.GetTableById(&table, string(blogId)); err != nil {
		ValidateError(c, DBError, err.Error())
		return
	}
	table.Title = req.Title
	table.Tag = strings.Join(req.Tag, ",")
	table.Desc = req.Desc
	table.Content = req.Content
	table.UpdateBy = c.Keys["user_id"].(uint32)
	if err := s.dao.SaveTable(&table); err != nil {
		ValidateError(c, DBError, "修改失败")
		return
	}
	Response(c, "修改文章成功")
}

func (s *RijiServer) BlogDelete(c *gin.Context) {
	blogId := c.Param("id")
	if blogId == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	var table model.Blog
	if err := s.dao.GetTableById(&table, blogId); err != nil {
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

func (s *RijiServer) BlogDetail(c *gin.Context) {
	blogId := c.Param("id")
	if blogId == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	var table model.Blog
	if err := s.dao.GetTableById(&table, blogId); err != nil {
		ValidateError(c, DBError, err.Error())
		return
	}
	data := map[string]interface{}{
		"id":      table.ID,
		"title":   table.Title,
		"tag":     strings.Split(table.Tag, ","),
		"desc":    table.Desc,
		"content": table.Content,
		"date":    table.CreatedAt.Format("2006-01-02"),
		"user":    s.dao.GetUserName(table.AddBy),
	}
	Response(c, data)
}
