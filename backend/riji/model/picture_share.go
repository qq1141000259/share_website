package model

// 背景图片表
type BackImgCos struct {
	Base
	Suffix string `json:"suffix" gorm:"column:c_suffix;size:8"`
}

// 留言记录表
type MessageBoard struct {
	Base
	Content string `json:"content" gorm:"column:c_content;type:text"`
}

// 用户上传数据表
type UserUploadFile struct {
	Base
	Title      string     `json:"title" gorm:"column:c_title;size:32"`
	Filename   string     `json:"filename" gorm:"column:c_filename;size:128"`
	FileType   string     `json:"filetype" gorm:"column:c_filetype;size:8"`
	Desc       string     `json:"desc" gorm:"column:c_desc;size:256"`
	OriginFile string     `json:"origin_file" gorm:"column:c_origin_file;size:32"`
	UploadType UploadType `json:"upload_type" gorm:"column:c_upload_type;size:8"`
}
