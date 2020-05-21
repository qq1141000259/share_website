package service

import (
	"github.com/gin-gonic/gin"
	"riji/model"
	. "riji/utils"
	"strconv"
)

type WordBookPostReq struct {
	Word       string `json:"word"`
	Annotation string `json:"annotation"`
	Remark     string `json:"remark"`
}

func (s *RijiServer) WordBookGet(c *gin.Context) {
	p, limit := PaginationHandler(c.Query("p"), c.Query("limit"))
	var tables []*model.WordBook
	var count int
	conditions := make(map[string]interface{})
	if c.Query("word") != "" {
		conditions["c_word"] = c.Query("word")
	}
	if c.Query("annotation") != "" {
		conditions["c_annotation"] = "%" + c.Query("annotation") + "%"
	}
	if c.Query("remark") != "" {
		conditions["c_remark"] = "%" + c.Query("remark") + "%"
	}
	if err := s.dao.GetTablePagination(&tables, &count, p, limit, conditions); err != nil {
		ValidateError(c, DBError, "查询单词记录失败")
		return
	}
	var data []map[string]interface{}
	for _, obj := range tables {
		data = append(data, map[string]interface{}{
			"id":         obj.ID,
			"word":       obj.Word,
			"annotation": obj.Annotation,
			"remark":     obj.Remark,
			"date":       obj.CreatedAt.Format("2006-01-02"),
			"user":       s.dao.GetUserName(obj.AddBy),
		})
	}
	Response(c, map[string]interface{}{
		"rows":  data,
		"total": count,
	})
}

func (s *RijiServer) WordBookPost(c *gin.Context) {
	var req WordBookPostReq
	if err := c.BindJSON(&req); err != nil {
		ValidateError(c, ParamError, "请求参数格式错误")
		return
	}
	if req.Word == "" || req.Annotation == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	table := model.WordBook{
		Annotation: req.Annotation,
		Word:       req.Word,
		Remark:     req.Remark,
	}
	table.AddBy = c.Keys["user_id"].(uint32)
	if err := s.dao.CreateTable(&table); err != nil {
		ValidateError(c, DBError, "插入数据失败")
		return
	}
	Response(c, "新增成功")
}

func (s *RijiServer) WordBookPut(c *gin.Context) {
	var req WordBookPostReq
	wordId := c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		ValidateError(c, ParamError, "请求参数格式错误")
		return
	}
	if wordId == "" || req.Word == "" || req.Annotation == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	id, err := strconv.Atoi(wordId)
	if err != nil{
		ValidateError(c, ParamError, "id必须是整数")
		return
	}
	table := model.WordBook{}
	if err := s.dao.GetWordBookById(&table, uint32(id)); err != nil {
		ValidateError(c, DBError, err.Error())
		return
	}
	table.Word = req.Word
	table.Annotation = req.Annotation
	table.Remark = req.Remark
	table.UpdateBy = c.Keys["user_id"].(uint32)
	if err := s.dao.SaveTable(&table); err != nil {
		ValidateError(c, DBError, "修改数据失败")
		return
	}
	Response(c, "修改成功")
}

func (s *RijiServer) WordBookDelete(c *gin.Context) {
	wordId := c.Param("id")
	if wordId == "" {
		ValidateError(c, ParamError, "缺少参数")
		return
	}
	id, err := strconv.Atoi(wordId)
	if err != nil{
		ValidateError(c, ParamError, "id必须是整数")
		return
	}
	var table model.WordBook
	if err := s.dao.GetWordBookById(&table, uint32(id)); err != nil {
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
