package services

import (
	"errors"
	"fmt"
	"my_super_project/models"
)

// мок (заглушка) - вместо БД
var items = []models.Item{
	{Id: 1, Name: "Laptop", Price: 1000.50, Amount: 10},
	{Id: 2, Name: "Phone", Price: 500.75, Amount: 5},
}

func GetAllItems() []models.Item {
	return items
}

func GetItemById(id int) (models.Item, error) {
	for _, v := range items {
		if v.Id == id {
			return v, nil
		}
	}
	return models.Item{}, fmt.Errorf("item with id = %d not found", id)
}

func AddItem(item models.Item) error {
	for _, v := range items {
		if v.Name == item.Name {
			return fmt.Errorf("item with name %s already exists", item.Name)
		}
	}

	item.Id = len(items) + 1
	items = append(items, item)

	return nil
}

func BuyItem(item models.Item) error {
	for i, v := range items {
		if v.Id == item.Id {
			if v.Amount > 0 {
				items[i].Amount -= 1
				return nil
			}
		}
	}

	return errors.New("item not found")
}

func DeleteItem(id int) error {
	for i, item := range items {
		if item.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

func UpdateItem(id int, updatedItem models.UpdatedItem) (*models.Item, error) {
	for i, item := range items {
		if item.Id == id {
			items[i].Name = updatedItem.Name
			items[i].Amount = updatedItem.Amount
			items[i].Price = updatedItem.Price

			return &items[i], nil
		}
	}

	return nil, errors.New("item not found")
}
