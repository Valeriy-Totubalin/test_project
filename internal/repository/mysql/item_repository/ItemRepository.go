package item_repository

import (
	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/repository_interfaces"
	"github.com/Valeriy-Totubalin/test_project/internal/domain"
	"gorm.io/gorm"
)

type ItemrRepository struct {
	Gorm repository_interfaces.GetterGormDB
}

func NewItemrRepository(gorm repository_interfaces.GetterGormDB) *ItemrRepository {
	return &ItemrRepository{
		Gorm: gorm,
	}
}

func (repo *ItemrRepository) Create(item *domain.Item) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	result := db.Create(&Item{
		Name:   item.Name,
		UserId: item.UserId,
	})

	if nil != result.Error {
		return result.Error
	}

	return nil
}

func (repo *ItemrRepository) DeleteById(itemId int) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	result := db.Delete(&Item{
		Id: itemId,
	})

	if nil != result.Error {
		return result.Error
	}

	return nil
}

func (repo *ItemrRepository) GetAll() ([]*domain.Item, error) {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return nil, err
	}

	var items []*Item

	result := db.Limit(500).Find(&items)
	if nil != result.Error {
		return nil, result.Error
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

func (repo *ItemrRepository) Transfer(itemId int, userId int) error {
	db, err := repo.Gorm.GetDB()
	if nil != err {
		return err
	}

	item := Item{}
	err = db.First(&item, itemId).Error
	if nil != err {
		return err
	}

	item.UserId = userId

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&Item{Id: itemId}).Error; err != nil {
			return err
		}

		if err := tx.Create(item).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}
