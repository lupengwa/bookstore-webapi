package order

import (
	"bookstore-webapi/internal/api/order/entity"
	"bookstore-webapi/internal/platform/db"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Repository interface {
	SaveOrder(orderEntity entity.OrderEntity) (entity.OrderEntity, error)
	SaveOrderItems(orderEntity []entity.OrderItemEntity) error
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(ds *db.DataSource) *RepositoryImpl {
	if ds == nil {
		log.Panic("DB connection can't be nil")
	}
	return &RepositoryImpl{
		db: ds.Connection,
	}
}

func (repo *RepositoryImpl) SaveOrder(orderEntity entity.OrderEntity) (entity.OrderEntity, error) {
	result := repo.db.Clauses(clause.Returning{}).Create(&orderEntity)
	if result.Error != nil {
		return orderEntity, fmt.Errorf("SaveOrder: %w", result.Error)
	}
	return orderEntity, nil
}

func (repo *RepositoryImpl) SaveOrderItems(orderItems []entity.OrderItemEntity) error {
	result := repo.db.Create(&orderItems)
	if result.Error != nil {
		return fmt.Errorf("SaveOrderItems: %w", result.Error)
	}
	return nil
}
