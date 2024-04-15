package domain

type Shelf struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Product  *Product `json:"product"`
	Quantity int64    `json:"quantity"`
}
