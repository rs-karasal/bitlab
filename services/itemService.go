package services

import (
	"my_super_project/models"
	"my_super_project/repositories"
)

func GetAllItems() ([]models.Item, error) {
	return repositories.GetAllItems()
}

func GetItemById(id int) (models.Item, error) {
	item, err := repositories.GetItemById(id)
	if err != nil {
		return item, err
	}

	item.Id = id

	return item, err
}

// func AddItem(item models.Item) error {
// 	for _, v := range items {
// 		if v.Name == item.Name {
// 			return fmt.Errorf("item with name %s already exists", item.Name)
// 		}
// 	}

// 	item.Id = len(items) + 1
// 	items = append(items, item)

// 	return nil
// }

// func BuyItem(item models.Item) error {
// 	for i, v := range items {
// 		if v.Id == item.Id {
// 			if v.Amount > 0 {
// 				items[i].Amount -= 1
// 				return nil
// 			}
// 		}
// 	}

// 	return errors.New("item not found")
// }

// func DeleteItem(id int) error {
// 	for i, item := range items {
// 		if item.Id == id {
// 			items = append(items[:i], items[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return errors.New("item not found")
// }

// func UpdateItem(id int, updatedItem models.UpdatedItem) (*models.Item, error) {
// 	for i, item := range items {
// 		if item.Id == id {
// 			items[i].Name = updatedItem.Name
// 			items[i].Amount = updatedItem.Amount
// 			items[i].Price = updatedItem.Price

// 			return &items[i], nil
// 		}
// 	}

// 	return nil, errors.New("item not found")
// }
