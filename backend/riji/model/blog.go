package model

// 博客信息表
type Blog struct {
	Base
	Title   string `json:"title" gorm:"column:c_title;size:32"`
	Desc    string `json:"desc" gorm:"column:c_desc;size:255"`
	Tag     string `json:"tag" gorm:"column:c_tag;size:255"`
	Content string `json:"content" gorm:"column:c_content;type:text"`
}
