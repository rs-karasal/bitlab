package server

import (
	"my_super_project/config"
	"my_super_project/handlers"
	"my_super_project/middleware"
	"my_super_project/utils/logger"
	"net/http"
)

// start server
func Run(cfg *config.Config) {
	mux := http.NewServeMux() // мультиплексор (маршрутизатор)

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetItemsHandler), cfg).ServeHTTP(w, r)
			// middleware.BasicAuthMiddleware(handlers.GetItemsHandler).ServeHTTP(w, r)
			// } else if r.Method == http.MethodPost {
			// 	handlers.AddItemHandler(w, r)
			// } else if r.Method == http.MethodDelete {
			// 	handlers.DeleteItemHandler(w, r)
			// } else if r.Method == http.MethodPut {
			// 	handlers.UpdateItemHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// mux.HandleFunc("/items/buy", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodPost {
	// 		handlers.BuyHandler(w, r)
	// 	} else {
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.RegisterHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/users/auth", handlers.AuthHandler(cfg))

	logger.InfoLog.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.ErrorLog.Fatalf("Error starting server: %v", err)
	}
}
