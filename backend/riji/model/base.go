package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UploadType string

const (
	PhotoType UploadType = "photo"
	VideoType UploadType = "video"
)

type Base struct {
	ID        string    `json:"id" gorm:"column:c_id;primary_key"`
	CreatedAt time.Time `json:"add_dt" gorm:"column:c_add_dt"`
	AddBy     uint32    `json:"add_by" gorm:"column:c_add_by"`
	UpdatedAt time.Time `json:"update_dt" gorm:"column:c_update_dt"`
	UpdateBy  uint32    `json:"update_by" gorm:"column:c_update_by"`
	IsDelete  bool      `json:"is_delete" gorm:"column:c_is_delete"`
}

// 将自增数字id改为uuid
func (user *Base) BeforeCreate(scope *gorm.Scope) error {
	if scope.TableName() == "t_user_upload_file" {
		return nil
	}
	newId := strings.Replace(uuid.New().String(), "-", "", -1)
	scope.SetColumn("ID", newId)
	return nil
}

type IntBase struct {
	ID        int    	`json:"id" gorm:"column:c_id;primary_key"`
	CreatedAt time.Time `json:"add_dt" gorm:"column:c_add_dt"`
	AddBy     uint32    `json:"add_by" gorm:"column:c_add_by"`
	UpdatedAt time.Time `json:"update_dt" gorm:"column:c_update_dt"`
	UpdateBy  uint32    `json:"update_by" gorm:"column:c_update_by"`
	IsDelete  bool      `json:"is_delete" gorm:"column:c_is_delete"`
}
