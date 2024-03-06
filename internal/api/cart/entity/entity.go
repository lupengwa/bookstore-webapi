package entity

import (
	"time"
)

type UserEntity struct {
	ID        string    `gorm:"primaryKey;column:id;"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (UserEntity) TableName() string {
	return "bookstore.users"
}

type CartItemEntity struct {
	SkuId     string    `gorm:"primaryKey;column:sku_id"`
	CartId    string    `gorm:"primaryKey;column:cart_id"`
	Quantity  int       `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (CartItemEntity) TableName() string {
	return "bookstore.cart_items"
}

type CartEntity struct {
	ID        string    `gorm:"primaryKey;column:id"`
	Total     float64   `gorm:"column:total"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (CartEntity) TableName() string {
	return "bookstore.cart"
}
