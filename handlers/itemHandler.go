package handlers

import (
	"encoding/json"
	"fmt"
	"my_super_project/models"
	"my_super_project/services"
	"net/http"
)

// обработка запроса для получения всех товаров (/items)
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := services.GetAllItems()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	services.AddItem(newItem)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Item added successfully")
}
