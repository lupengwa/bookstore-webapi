package mapper

import (
	"bookstore-webapi/internal/api/cart/dto"
	"bookstore-webapi/internal/api/cart/entity"
)

func CartItemEntityToCartItemDto(entity entity.CartItemEntity) dto.CartItemDto {
	return dto.CartItemDto{
		SkuId:    entity.SkuId,
		Quantity: entity.Quantity,
	}
}

func CartItemDtoToCartItemEntity(dto dto.CartItemDto, cartId string) entity.CartItemEntity {
	return entity.CartItemEntity{
		CartId:   cartId,
		SkuId:    dto.SkuId,
		Quantity: dto.Quantity,
	}
}

func CartDtoToCartEntity(cartDto dto.CartDto) entity.CartEntity {
	return entity.CartEntity{
		ID:    cartDto.CartId,
		Total: cartDto.Total,
	}

}
