package orm

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id     int
	Name   string
	UserId int
}
