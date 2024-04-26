package domain

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Sku      string `json:"sku"`
	Category string `json:"category"`
}
