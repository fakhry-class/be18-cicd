package service

import "fakhry/cleanarch/features/item"

type itemService struct {
	itemData item.ItemDataInterface
}

// GetAll implements item.ItemServiceInterface.
func (service *itemService) GetAll(name string) ([]item.Core, error) {
	// if name != "" {
	// 	service.itemData.SelectByName(name)
	// }
	return service.itemData.SelectAll(name)
}

func New(repo item.ItemDataInterface) item.ItemServiceInterface {
	return &itemService{
		itemData: repo,
	}
}
