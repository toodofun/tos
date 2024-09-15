package database

import "gorm.io/gorm"

type BaseMapper[T any] struct {
	DB *Database
}

func NewMapper[T any](db *Database, model T) *BaseMapper[T] {
	return &BaseMapper[T]{
		DB: db,
	}
}

type Pager[T any] struct {
	CurrentPage int64 `json:"current"`
	PageSize    int64 `json:"size"`
	Total       int64 `json:"total"`
	Data        []T   `json:"data"`
}

func (m *BaseMapper[T]) Insert(entity T, db ...*gorm.DB) error {
	if len(db) > 0 {
		return db[0].Create(entity).Error
	}
	return m.DB.Create(entity).Error
}

func (m *BaseMapper[T]) Delete(entity T, db ...*gorm.DB) error {
	if len(db) > 0 {
		return db[0].Delete(entity).Error
	}
	return m.DB.Delete(entity).Error
}

func (m *BaseMapper[T]) Update(entity T, fields map[string]interface{}, db ...*gorm.DB) error {
	if len(db) > 0 {
		return db[0].Model(entity).Where(entity).Updates(fields).Error
	}
	return m.DB.Model(entity).Where(entity).Updates(fields).Error
}

func (m *BaseMapper[T]) Detail(entity T) (res T, err error) {
	err = m.DB.First(&res, entity).Error
	return
}

func (m *BaseMapper[T]) List(entity T, order ...string) (res []T, err error) {
	if len(order) > 0 {
		err = m.DB.Order(order).Find(&res, entity).Error
	} else {
		err = m.DB.Find(&res, entity).Error
	}
	return
}

func (m *BaseMapper[T]) Count(entity T) (count int64, err error) {
	err = m.DB.Model(entity).Where(entity).Count(&count).Error
	return
}

func (m *BaseMapper[T]) Page(entity T, page, size int64) (res *Pager[T], err error) {
	res = new(Pager[T])
	res.CurrentPage = page
	res.PageSize = size
	m.DB.Model(&entity).Where(entity).Count(&res.Total)
	if res.Total == 0 {
		res.Data = make([]T, 0)
	}
	err = m.DB.Model(&entity).Order("updated_at desc").Where(entity).Scopes(Pagination(res)).Find(&res.Data).Error
	return
}

func Pagination[T any](pager *Pager[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		size := pager.PageSize
		page := pager.CurrentPage
		offset := int((page - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
