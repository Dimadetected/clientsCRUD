package clientsCRUD

type Client struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Email string `json:"email" db:"name" binding:"required"`
	Age   int    `json:"age" db:"age" binding:"required"`
}
