package cart

import (
	"bookstore-webapi/internal/api/cart/dto"
	"bookstore-webapi/internal/api/cart/entity"
	"bookstore-webapi/internal/api/mapperutils"
	"bookstore-webapi/internal/mapper"
	"fmt"
	"log"
)

type Service interface {
	GetCartItems(cartId string) (dto.CartDto, error)
	ClearCart(cartId string) error
	AddCartItems(cartDto dto.CartDto, cartId string) (dto.CartDto, error)
	CreateNewCart(cartId string) (bool, dto.CartDto, error)
}

type ServiceImpl struct {
	cartRepo Repository
}

func NewService(cartRepo Repository) *ServiceImpl {
	if cartRepo == nil {
		log.Panic("Cart repo can't be nil")
	}
	return &ServiceImpl{cartRepo: cartRepo}
}

func (service ServiceImpl) GetCartItems(cartId string) (dto.CartDto, error) {
	var cartDto dto.CartDto
	cartEntity, cartItemsEntity, err := service.cartRepo.FindCartAndItemsByCartId(cartId)
	if err != nil {
		return cartDto, fmt.Errorf("GetCartItems: %w\n", err)
	}
	result := mapperutils.EntityListToDtoList[entity.CartItemEntity, dto.CartItemDto](cartItemsEntity, mapper.CartItemEntityToCartItemDto)
	cartDto = dto.CartDto{CartId: cartId, Total: cartEntity.Total, CartItems: result}
	return cartDto, nil
}

func (service ServiceImpl) ClearCart(cartId string) error {
	err := service.cartRepo.DeleteCartItemsByCartId(cartId)
	if err != nil {
		return fmt.Errorf("ClearCart: %w", err)
	}
	return nil
}

func (service ServiceImpl) AddCartItems(cartDto dto.CartDto, cartId string) (dto.CartDto, error) {
	cartEntity := mapper.CartDtoToCartEntity(cartDto)
	cartItemsEntity := mapperutils.DtoListToEntityList[dto.CartItemDto, entity.CartItemEntity](cartDto.CartItems, cartId, mapper.CartItemDtoToCartItemEntity)
	var cartUpdated dto.CartDto
	if err := service.cartRepo.SaveCartAndItems(cartItemsEntity, cartEntity); err != nil {
		return cartUpdated, fmt.Errorf("AddCartItems: %w", err)
	} else {
		cartEntity, cartItemsEntity, err := service.cartRepo.FindCartAndItemsByCartId(cartId)
		if err != nil {
			return cartUpdated, fmt.Errorf("AddCartItems: %w", err)
		}
		result := mapperutils.EntityListToDtoList[entity.CartItemEntity, dto.CartItemDto](cartItemsEntity, mapper.CartItemEntityToCartItemDto)
		cartUpdated = dto.NewCartDto(result, cartId, cartEntity.Total)
		return cartUpdated, nil
	}
}

// CreateNewCart
// return true and new cart if new cart created
// return false and existing cart if cart already exists
func (service ServiceImpl) CreateNewCart(cartId string) (bool, dto.CartDto, error) {
	cartDto := dto.CartDto{CartId: cartId}

	// check if cart already exists
	cartEntity, err := service.cartRepo.FindCartById(cartId)
	if err != nil {
		return false, cartDto, fmt.Errorf("CreateNewCart: %w\n", err)
	}

	// carEntity not found, create new cart
	if cartEntity.ID == "" {
		cartEntity = entity.CartEntity{
			ID:    cartId,
			Total: 0.00,
		}
		if err = service.cartRepo.SaveCart(cartEntity); err != nil {
			return false, cartDto, fmt.Errorf("CreateNewCart: failed to save new cart %w\n", err)
		}
		return true, cartDto, nil
	} else {
		_, cartItemsEntity, err := service.cartRepo.FindCartAndItemsByCartId(cartId)
		if err != nil {
			return false, cartDto, fmt.Errorf("CreateNewCart: failed to save new cart %w\n", err)
		}
		result := mapperutils.EntityListToDtoList[entity.CartItemEntity, dto.CartItemDto](cartItemsEntity, mapper.CartItemEntityToCartItemDto)
		cartDto.CartItems = result
		cartDto.Total = cartEntity.Total
		return false, cartDto, nil
	}

}
