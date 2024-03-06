package cart

import (
	"bookstore-webapi/internal/api/cart/entity"
	"bookstore-webapi/internal/platform/db"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	SaveCartAndItems(cartItems []entity.CartItemEntity, cartEntity entity.CartEntity) error
	SaveCart(cart entity.CartEntity) error
	DeleteByCartId(cartId string) error
	DeleteCartItemsByCartId(cartId string) error
	FindCartById(cartId string) (entity.CartEntity, error)
	FindCartAndItemsByCartId(cartId string) (entity.CartEntity, []entity.CartItemEntity, error)
	FindItemsByCartIdAndSkuId(cartId string, skuId string) ([]entity.CartItemEntity, error)
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

func (repo *RepositoryImpl) SaveCart(cart entity.CartEntity) (err error) {
	result := repo.db.Create(&cart)
	if result.Error != nil {
		return fmt.Errorf("SaveCart: %w", result.Error)
	}
	return nil
}

func (repo *RepositoryImpl) SaveCartAndItems(cartItems []entity.CartItemEntity, cartEntity entity.CartEntity) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Updates(&cartEntity).Error; err != nil {
			return err
		}
		if err := tx.Create(&cartItems).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return fmt.Errorf("SaveCartAndItems: %w", err)
	}
	return nil
}

func (repo *RepositoryImpl) FindCartById(cartId string) (entity.CartEntity, error) {
	var cartEntity entity.CartEntity
	err := repo.db.Where("id = ?", cartId).Find(&cartEntity).Error
	if err != nil {
		return cartEntity, fmt.Errorf("FindCartById: %w", err)
	}
	return cartEntity, nil
}

func (repo *RepositoryImpl) DeleteByCartId(cartId string) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Where("cart_id = ?", cartId).Delete(&entity.CartItemEntity{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", cartId).Delete(&entity.CartEntity{}).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return fmt.Errorf("DeleteByCartId: %w", err)
	}
	return nil
}

func (repo *RepositoryImpl) DeleteCartItemsByCartId(cartId string) error {
	if err := repo.db.Where("cart_id = ?", cartId).Delete(&entity.CartItemEntity{}).Error; err != nil {
		return fmt.Errorf("DeleteCartItemsByCartId: %w", err)
	}
	return nil
}

func (repo *RepositoryImpl) FindCartAndItemsByCartId(cartId string) (entity.CartEntity, []entity.CartItemEntity, error) {
	var cartItems []entity.CartItemEntity
	var cartEntity entity.CartEntity
	err := repo.db.Where("cart_id = ?", cartId).Find(&cartItems).Error
	if err != nil {
		return cartEntity, cartItems, fmt.Errorf("FindCartAndItemsByCartId: %w", err)
	}
	err = repo.db.Where("id = ?", cartId).Find(&cartEntity).Error
	if err != nil {
		return cartEntity, cartItems, fmt.Errorf("FindCartAndItemsByCartId: %w", err)
	}
	return cartEntity, cartItems, nil
}

func (repo *RepositoryImpl) FindItemsByCartIdAndSkuId(cartId string, skuId string) ([]entity.CartItemEntity, error) {
	var cartItems []entity.CartItemEntity
	err := repo.db.Where("id = ? and sku_id = ? ", cartId, skuId).Find(&cartItems).Error
	if err != nil {
		return cartItems, fmt.Errorf("FindItemsByCartIdAndSkuId: %w", err)
	}
	return cartItems, nil
}
