package services

import "my_super_project/models"

// мок (заглушка) - вместо БД
var items = []models.Item{
	{Id: 1, Name: "Laptop", Price: 1000.50, Amount: 10},
	{Id: 2, Name: "Phone", Price: 500.75, Amount: 5},
}

func GetAllItems() []models.Item {
	return items
}

func AddItem(item models.Item) {
	item.Id = len(items) + 1
	items = append(items, item)
}
