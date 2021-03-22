package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int
	Login    string
	Password string
}
