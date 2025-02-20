package handlers

import (
	"encoding/json"
	"my_super_project/models"
	"my_super_project/services"
	"my_super_project/utils/logger"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.ErrorLog.Print("Invalid input: " + err.Error())
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.RegisterService(user)
	if err != nil {
		logger.ErrorLog.Printf("Error registering user: %v", err)
		http.Error(w, "Error registering new user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"Status": "Success"})
}
