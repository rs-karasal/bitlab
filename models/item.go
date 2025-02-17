package models

type Item struct {
	Id     int     `json:"id"` // tags ``
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

type UpdatedItem struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}
