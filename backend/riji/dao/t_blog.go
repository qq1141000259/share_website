package dao

import "riji/model"

// 查处所有
func (d *Dao) ListBlogInfo(table *[]*model.MessageBoard) error {
	return d.Db.Find(table).Error
}
