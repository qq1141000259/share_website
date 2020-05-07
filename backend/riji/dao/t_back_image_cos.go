package dao

import "riji/model"

// 查处所有
func (d *Dao) ListBackImageCos(table *[]*model.BackImgCos) error {
	return d.Db.Find(table).Error
}

// 查询id
func (d *Dao) GetBackImageCosById(table *model.BackImgCos) error {
	return d.Db.Where("id = ?", table.ID).First(table).Error
}

// 分页查询
func (d *Dao) BackImageCosPagination(table *[]*model.BackImgCos, count *uint32, page,perPage uint32) error {
	return d.Db.Model(table).Count(count).Limit(perPage).Offset((page - 1) * perPage).Order("c_add_dt desc").Find(table).Error
}

// 新增记录
func (d *Dao) CreateBackImageCos(table *model.BackImgCos) error {
	return d.Db.Create(table).Error
}

// 修改记录方式1，直接在原纪录上修改
func (d *Dao) SaveBackImageCos(table *model.BackImgCos) error {
	return d.Db.Save(table).Error
}

// 传一个字典格式的修改
func (d *Dao) UpdateBackImageCos(table *model.BackImgCos, attrs interface{}) error {
	return d.Db.Model(table).Update(attrs).Error
}
