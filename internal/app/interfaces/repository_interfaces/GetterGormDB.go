package repository_interfaces

import "gorm.io/gorm"

type GetterGormDB interface {
	GetDB() (*gorm.DB, error)
}
