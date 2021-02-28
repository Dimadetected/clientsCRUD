package clientsCRUD

type Sale struct {
	Id    int     `json:"-" db:"id"`
	Count int     `json:"count" binding:"required" db:"count"`
	Total float64 `json:"total" binding:"required" db:"total"`
}

type ClientSale struct {
	Id       int `json:"-" db:"id"`
	ClientId int `json:"client_id" binding:"required" db:"client_id"`
	SaleId   int `json:"sale_id" binding:"required" db:"sale_id"`
}
