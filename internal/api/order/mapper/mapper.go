package mapper

import (
	"bookstore-webapi/internal/api/order/dto"
	"bookstore-webapi/internal/api/order/entity"
)

func OrderItemEntityToOrderItemDto(orderItem entity.OrderItemEntity) dto.OrderItemDto {
	return dto.OrderItemDto{
		SkuId:    orderItem.SkuId,
		Quantity: orderItem.Quantity,
	}
}
