package item_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/test_project/db/orm"
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"gorm.io/gorm"
)

type ItemRepository struct {
	Gorm repository_interfaces.GetterGormDB
}

func NewItemRepository(gorm repository_interfaces.GetterGormDB) *ItemRepository {
	return &ItemRepository{
		Gorm: gorm,
	}
}

func (repo *ItemRepository) Create(item *domain.Item) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	itemOrm := &orm.Item{
		Name:   item.Name,
		UserId: item.UserId,
	}

	err = db.Create(itemOrm).Error
	if nil != err {
		return err
	}
	item.Id = itemOrm.Id

	return nil
}

func (repo *ItemRepository) DeleteById(itemId int) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	err = db.Delete(&orm.Item{
		Id: itemId,
	}).Error

	if nil != err {
		return err
	}

	return nil
}

func (repo *ItemRepository) GetAll(userId int) ([]*domain.Item, error) {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return nil, err
	}

	var items []*orm.Item

	err = db.Limit(500).Find(&items).Error
	if nil != err {
		return nil, err
	}

	var domainItems []*domain.Item
	for _, item := range items {
		domainItems = append(domainItems, &domain.Item{
			Id:     item.Id,
			Name:   item.Name,
			UserId: item.UserId,
		})
	}

	return domainItems, nil
}

func (repo *ItemRepository) Transfer(itemId int, userId int) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	item := orm.Item{}
	err = db.First(&item, itemId).Error
	if nil != err {
		return err
	}

	item.UserId = userId

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&orm.Item{Id: itemId}).Error; err != nil {
			return err
		}

		if err := tx.Create(&orm.Item{Name: item.Name, UserId: item.UserId}).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (repo *ItemRepository) GetById(itemId int) (*domain.Item, error) {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return nil, err
	}

	item := orm.Item{}
	err = db.Find(&item, itemId).Error
	if nil != err {
		return nil, err
	}

	if 0 == item.Id {
		return nil, errors.New("item does not exist")
	}

	return &domain.Item{
		Id:     item.Id,
		Name:   item.Name,
		UserId: item.UserId,
	}, nil
}
