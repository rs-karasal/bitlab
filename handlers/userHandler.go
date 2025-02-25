package handlers

import (
	"encoding/json"
	"my_super_project/config"
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

func AuthHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var userAuth models.UserDTO
			err := json.NewDecoder(r.Body).Decode(&userAuth)
			if err != nil {
				logger.ErrorLog.Print("Invalid input: " + err.Error())
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			token, err := services.AuthService(userAuth, cfg)
			if err != nil {
				logger.ErrorLog.Printf("Error login user: %v", err)
				http.Error(w, "Error logging user", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if token == "" {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"Status": "Error", "token": token})
			} else {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(map[string]string{"Status": "Success", "token": token})
			}

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	}
}
