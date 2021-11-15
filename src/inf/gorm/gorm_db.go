package gorm

import "gorm.io/gorm"

type DbGorm interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)

	Create(value interface{}) (tx *gorm.DB)

	Updates(values interface{}) (tx *gorm.DB)

	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)

	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}
