package orm

import (
	"errors"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDB struct {
	Config config_interfaces.DBConfig
}

func NewGormDB(config config_interfaces.DBConfig) *GormDB {
	return &GormDB{
		Config: config,
	}
}

func (g *GormDB) GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(g.generateDSN()), &gorm.Config{})
	if err != nil {
		return nil, errors.New("error connecting to database")
	}

	return db, nil
}

func (g *GormDB) generateDSN() string {
	return g.Config.GetUser() + ":" +
		g.Config.GetPassword() + "@tcp(" +
		g.Config.GetHost() + ":" +
		g.Config.GetPort() + ")/" +
		g.Config.GetName() + "?charset=utf8mb4&parseTime=True&loc=Local"
}
