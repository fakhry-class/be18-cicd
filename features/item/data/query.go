package data

import (
	"fakhry/cleanarch/features/item"

	"gorm.io/gorm"
)

type itemQuery struct {
	db *gorm.DB
}

// SelectAll implements item.ItemDataInterface.
func (repo *itemQuery) SelectAll(name string) ([]item.Core, error) {
	var itemsData []Item
	var tx *gorm.DB
	if name != "" {
		tx = repo.db.Preload("User").Where("name like ?", "%"+name+"%").Find(&itemsData) // select * from item;
	} else {
		tx = repo.db.Preload("User").Find(&itemsData) // select * from item;
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	// fmt.Println("items:", itemsData)
	//mapping dari struct gorm model ke struct core (entity)
	var itemsCore []item.Core
	for _, value := range itemsData {
		var user = item.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Brand:       value.Brand,
			Description: value.Description,
			Price:       value.Price,
			Weight:      value.Weight,
			User: item.UserCore{
				ID:    value.User.ID,
				Name:  value.User.Name,
				Email: value.User.Email,
			},
		}
		itemsCore = append(itemsCore, user)
	}
	return itemsCore, nil
}

func New(db *gorm.DB) item.ItemDataInterface {
	return &itemQuery{
		db: db,
	}
}
