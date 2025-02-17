package handlers

import (
	"encoding/json"
	"fmt"
	"my_super_project/services"
	"net/http"
	"strconv"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid Id parameter", http.StatusBadRequest)
			return
		}

		item, err := services.GetItemById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
		return
	}

	items, err := services.GetAllItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch items: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// func AddItemHandler(w http.ResponseWriter, r *http.Request) {
// 	var newItem models.Item

// 	err := json.NewDecoder(r.Body).Decode(&newItem)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	if err = services.AddItem(newItem); err != nil {
// 		log.Printf("Can not add new item: %v", err)
// 		http.Error(w, err.Error(), http.StatusConflict)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "Item added successfully")
// }

// func BuyHandler(w http.ResponseWriter, r *http.Request) {
// 	var buyItem models.Item

// 	err := json.NewDecoder(r.Body).Decode(&buyItem)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	err = services.BuyItem(buyItem)
// 	if err != nil {
// 		http.Error(w, "Item not found", http.StatusNotFound)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Enjoy your new item :)")
// }

// func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
// 	idStr := r.URL.Query().Get("id")
// 	if idStr != "" {
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			log.Printf("Error converting idStr to id: %v", err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		if err = services.DeleteItem(id); err != nil {
// 			log.Printf("Error deleting item: %v", err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte("Successfully deleted an item"))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusBadRequest)
// 	w.Write([]byte("Caould not delete item"))
// }

// func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
// 	idStr := r.URL.Query().Get("id")
// 	if idStr == "" {
// 		http.Error(w, "id is empty", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "wrong id format", http.StatusBadRequest)
// 		return
// 	}

// 	var updatedItem models.UpdatedItem
// 	err = json.NewDecoder(r.Body).Decode(&updatedItem)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	item, err := services.UpdateItem(id, updatedItem)
// 	if err != nil {
// 		if err.Error() == "item not found" {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, "failed to update item", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(item)
// }
