package clientsCRUD

type Client struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Email string `json:"email" db:"email" binding:"required"`
	Age   int    `json:"age" db:"age" binding:"required"`
}

type UpdateClient struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int    `json:"age"`
}
