package dao

import "riji/model"

func (d *Dao) ListWordBook(table *[]*model.WordBook) error {
	return d.Db.Find(table).Error
}

func (d *Dao) GetWordBookById(table *model.WordBook, id uint32) error {
	return d.Db.Where("c_id = ?", id).First(table).Error
}
