package order

import (
	"bookstore-webapi/internal/api/mapperutils"
	"bookstore-webapi/internal/api/order/dto"
	"bookstore-webapi/internal/api/order/entity"
	"bookstore-webapi/internal/api/order/mapper"
	"fmt"
	"log"
)

type Service interface {
	CreateNewOrder(userId string, total float64, orderItems []entity.OrderItemEntity) (dto.OrderDto, error)
}

type ServiceImpl struct {
	orderRepo Repository
}

func NewService(orderRepo Repository) *ServiceImpl {
	if orderRepo == nil {
		log.Panic("Cart repo can't be nil")
	}
	return &ServiceImpl{orderRepo: orderRepo}
}

func (service ServiceImpl) CreateNewOrder(userId string, total float64, orderItems []entity.OrderItemEntity) (dto.OrderDto, error) {
	var orderResp dto.OrderDto

	// save order
	orderToSave := entity.OrderEntity{
		UserId: userId,
		Total:  total,
	}
	orderSaved, err := service.orderRepo.SaveOrder(orderToSave)
	orderId := orderSaved.ID.String()
	if err != nil {
		return orderResp, fmt.Errorf("CreateNewOrder: %w", err)
	}
	// save order items
	for i, _ := range orderItems {
		orderItems[i].OrderId = orderId
	}
	err = service.orderRepo.SaveOrderItems(orderItems)
	if err != nil {
		return orderResp, fmt.Errorf("CreateNewOrder: %w", err)
	}

	// return order response
	orderItemsDto := mapperutils.EntityListToDtoList[entity.OrderItemEntity, dto.OrderItemDto](orderItems, mapper.OrderItemEntityToOrderItemDto)
	orderResp = dto.OrderDto{
		OrderId:    orderId,
		Total:      orderSaved.Total,
		OrderItems: orderItemsDto,
	}

	return orderResp, nil
}
