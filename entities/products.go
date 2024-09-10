package entities

import (
	"go-native/config"
	"time"
)

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	CreatedAt  time.Time
	UpdatedAt time.Time
}
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}