package dao

import "riji/model"

func (d *Dao) CheckUserName(name string) bool {
	if err := d.Db.Where("c_username = ?", name).First(&model.User{}).Error; err != nil{
		return true
	}
	return false
}


func (d *Dao) GetUserByUserName(table *model.User) error {
	return d.Db.Where("c_username = ?", table.UserName).First(table).Error
}
