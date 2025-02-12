package models

type Item struct {
	Id     int     `json:"id"` // tags ``
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amoun"`
}

type UpdatedItem struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amoun"`
}
