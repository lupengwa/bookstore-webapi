package dto

type CartItemDto struct {
	SkuId    string `json:"skuId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gte=1"`
}

type CartDto struct {
	CartId    string        `json:"cartId,omitempty"`
	Total     float64       `json:"total" validate:"required,gte=0""`
	CartItems []CartItemDto `json:"cartItems" validate:"dive"`
}

func NewCartDto(cartItems []CartItemDto, cartId string, total float64) CartDto {
	return CartDto{
		CartId:    cartId,
		Total:     total,
		CartItems: cartItems,
	}
}
