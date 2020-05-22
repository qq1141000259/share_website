package dao

import (
	"fmt"
	"reflect"
	"riji/config"
	"riji/model"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Dao 处理连接中间层
type Dao struct {
	Db *gorm.DB
}

// NewDao 新建 Dao
func NewDao(c *config.Config) (*Dao, error) {
	db, err := NewMySQL(c)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Dao{
		Db: db,
	}, err
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *config.Config) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", c.DSN)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return
}

// Ping pong
func (d *Dao) Ping() error {
	return d.Db.DB().Ping()
}

// Close close connection
func (d *Dao) Close() error {
	return d.Db.Close()
}

// 查处所有
func (d *Dao) ListTable(table interface{}) error {
	return d.Db.Find(table).Error
}

// 查询id
func (d *Dao) GetTableById(table interface{}, id string) error {
	return d.Db.Where("c_id = ?", id).First(table).Error
}

// 按创造时间获取
func (d *Dao) GetTableOrderByCreateTime(table interface{}) error {
	return d.Db.Model(table).Limit(5).Order("c_add_dt desc").Find(table).Error
}

// 分页查询
func (d *Dao) GetTablePagination(table interface{}, count *int, page, perPage int, conditions map[string]interface{}) error {
	param := []interface{}{false}
	sParam := "c_is_delete= ?"
	for k, v := range conditions {
		if reflect.TypeOf(v).String() == "string" && strings.HasSuffix(v.(string), "%"){
			sParam += fmt.Sprintf(" and %s like ?", k)
		} else {
			sParam += fmt.Sprintf(" and %s = ?", k)
		}
		param = append(param, v)
	}
	err := d.Db.Model(table).Limit(perPage).Offset((page-1)*perPage).Order("c_add_dt desc").Where(sParam, param...).Find(table).Error
	if err != nil {
		return err
	}
	err = d.Db.Model(table).Where(sParam, param...).Count(count).Error
	return err
}

// 新增记录
func (d *Dao) CreateTable(table interface{}) error {
	return d.Db.Create(table).Error
}

// 修改记录方式1，直接在原纪录上修改
func (d *Dao) SaveTable(table interface{}) error {
	return d.Db.Save(table).Error
}

// 传一个字典格式的修改
func (d *Dao) UpdateTable(table interface{}, attrs interface{}) error {
	return d.Db.Model(table).Update(attrs).Error
}

// 查询user
func (d *Dao) GetUserName(id uint32) string {
	var table model.User
	if err := d.Db.Where("c_id = ?", id).First(&table).Error; err != nil {
		return ""
	}
	return table.Name
}
