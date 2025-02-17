package repositories

import (
	"fmt"
	"my_super_project/database"
	"my_super_project/models"
)

func GetAllItems() ([]models.Item, error) {
	query := `
	SELECT id, name, price, amount
	FROM items
	`

	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return items, nil
}

func GetItemById(id int) (models.Item, error) {
	query := `
	SELECT name, price, amount
	FROM items
	WHERE id = $1;
	`

	var itemById models.Item
	err := database.Db.QueryRow(query, id).Scan(&itemById.Name, &itemById.Price, &itemById.Amount)
	if err != nil {
		return models.Item{}, fmt.Errorf("failed to retrieve item with ID %d: %v", id, err)
	}

	return itemById, nil
}
