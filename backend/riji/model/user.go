package model

import (
	"time"
)

type User struct {
	ID        uint32    `json:"id" gorm:"column:c_id;primary_key"`
	CreatedAt time.Time `json:"add_dt" gorm:"column:c_add_dt"`
	UpdatedAt time.Time `json:"update_dt" gorm:"column:c_update_dt"`
	Email     string    `json:"email" gorm:"column:c_email;size:64"`
	Name      string    `json:"name" gorm:"column:c_name;size:64"`
	UserName  string    `json:"username" gorm:"column:c_username;size:64;unique"`
	Password  string    `json:"password" gorm:"column:c_password;size:64"`
	// 想要初始化为null 可以设置为指针来传入nil
	LastLoginTime *time.Time `json:"last_login_time" gorm:"column:c_last_login_time;null"`
	LastLoginIp   string     `json:"last_login_ip" gorm:"column:c_last_login_ip;size:15"`
}
