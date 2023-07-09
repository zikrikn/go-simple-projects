package models

type Product struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:varchar(100)"`
	Price       int    `json:"price" gorm:"type:integer"`
}
