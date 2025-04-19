package dto

type CreateItemIput struct {
	Name string `json:"name" binding:"required,min=2"`
	Price uint `json:"price" binging:"required,min=1,max=999999"`
	Description string `json:"description"`
}
