package server

import (
	"log"
	"my_super_project/handlers"
	"net/http"
)

// start server
func Run() {
	mux := http.NewServeMux() // мультиплексор (маршрутизатор)

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetItemsHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.AddItemHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
