package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type OrderEntity struct {
	ID        uuid.UUID `gorm:"primaryKey;column:id;type:uuid;default:gen_random_uuid()"`
	UserId    string    `gorm:"column:user_id"`
	Total     float64   `gorm:"column:total"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (OrderEntity) TableName() string {
	return "bookstore.orders"
}

type OrderItemEntity struct {
	OrderId   string    `gorm:"primaryKey;column:order_id"`
	SkuId     string    `gorm:"primaryKey;column:sku_id"`
	Quantity  int       `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (OrderItemEntity) TableName() string {
	return "bookstore.order_items"
}
