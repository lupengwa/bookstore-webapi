package dto

type OrderItemDto struct {
	SkuId    string `json:"skuId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gte=1"`
}

type OrderDto struct {
	OrderId    string         `json:"orderId,omitempty"`
	OrderItems []OrderItemDto `json:"orderItems" validate:"dive"`
	Total      float64        `json:"total" validate:"required"`
}

func NewOrderDto(orderItems []OrderItemDto, orderId string) OrderDto {
	return OrderDto{
		OrderId:    orderId,
		OrderItems: orderItems,
	}
}
