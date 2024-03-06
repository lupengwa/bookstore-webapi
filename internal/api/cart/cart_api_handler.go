package cart

import (
	"bookstore-webapi/internal/api/cart/dto"
	"bookstore-webapi/internal/api/order"
	"bookstore-webapi/internal/api/order/entity"
	"bookstore-webapi/internal/api/restutils"
	"bookstore-webapi/internal/apperror"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

const (
	CartApiBasePath     = "/cart"
	CartApiResourcePath = CartApiBasePath + "/{cartId}"
	CartApiCheckoutPath = CartApiBasePath + "/{cartId}/checkout"
)

// ApiHandler is the API request handler for the cart domain
type ApiHandler struct {
	cartService  Service
	orderService order.Service
}

func NewApiHandler(cartService Service, orderService order.Service) *ApiHandler {
	if cartService == nil {
		log.Panic("cart service can't be nil")
	}
	if orderService == nil {
		log.Panic("order service can't be nil")
	}
	return &ApiHandler{cartService: cartService, orderService: orderService}
}

func (handler *ApiHandler) GetRestUriToHandlerConfig() map[restutils.RestApiUriKey]http.HandlerFunc {
	return map[restutils.RestApiUriKey]http.HandlerFunc{
		restutils.RestApiUriKey{
			HttpMethod: http.MethodGet,
			Path:       CartApiResourcePath,
		}: handler.ListCartItems,
		restutils.RestApiUriKey{
			HttpMethod: http.MethodPost,
			Path:       CartApiResourcePath,
		}: handler.UpdateCart,
		restutils.RestApiUriKey{
			HttpMethod: http.MethodPost,
			Path:       CartApiBasePath,
		}: handler.CreateCart,
		restutils.RestApiUriKey{
			HttpMethod: http.MethodPost,
			Path:       CartApiCheckoutPath,
		}: handler.CheckoutCart,
	}
}

// ListCartItems fetch all cart items
func (handler *ApiHandler) ListCartItems(w http.ResponseWriter, r *http.Request) {
	// validate user
	_, err := restutils.ValidateUser(r)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusUnauthorized)
		return
	}
	// extract request parameter
	cartId := chi.URLParam(r, "cartId")
	if cartId == "" {
		restutils.ToErrorResponse(w, r, apperror.MissingCartIdErr, http.StatusBadRequest)
		return
	}

	cart, err := handler.cartService.GetCartItems(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("ListCartItems: %v\n", err)
		return
	}
	restutils.ToSuccessPayloadResponse(w, r, cart)
}

// UpdateCart update cart in db
func (handler *ApiHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	_, err := restutils.ValidateUser(r)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusUnauthorized)
		return
	}

	cartId := chi.URLParam(r, "cartId")
	if cartId == "" {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusBadRequest)
		return
	}

	var cartDto dto.CartDto
	if err := restutils.UnmarshalJSONRequest(r, &cartDto); err != nil {
		restutils.ToErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	if err := restutils.ValidateJSONRequest(cartDto); err != nil {
		restutils.ToErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	// clear existing cart
	err = handler.cartService.ClearCartItems(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("UpdateCart: %s failed cause: %s", cartId, err)
		return
	}

	// update new cart
	updatedCart, err := handler.cartService.AddCartItems(cartDto, cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("UpdateCart: failed to update cart item. cause: %v\n", err)
		return
	}
	restutils.ToSuccessPayloadResponse(w, r, updatedCart)

}

// CreateCart creates a new cart
func (handler *ApiHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	userId, err := restutils.ValidateUser(r)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusUnauthorized)
		return
	}
	cartId := userId

	// use Location to inform client the resource directory
	w.Header().Set("Location", fmt.Sprintf("/cart/%s", cartId))

	isNewCart, cartDto, err := handler.cartService.CreateNewCart(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("CreateCart: %v\n", err)
		return
	}

	if isNewCart {
		render.Status(r, http.StatusCreated)
		render.Respond(w, r, cartDto)
	} else {
		restutils.ToSuccessPayloadResponse(w, r, cartDto)
	}
}

// CheckoutCart checkout a cart and create an order
func (handler *ApiHandler) CheckoutCart(w http.ResponseWriter, r *http.Request) {
	userId, err := restutils.ValidateUser(r)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusUnauthorized)
		return
	}

	cartId := chi.URLParam(r, "cartId")
	if cartId == "" {
		restutils.ToErrorResponse(w, r, apperror.MissingCartIdErr, http.StatusBadRequest)
		return
	}

	// assume payment is always successful
	paymentSuccess := true
	if !paymentSuccess {
		restutils.ToErrorResponse(w, r, apperror.PaymentErr, http.StatusInternalServerError)
		return
	}

	// fetch cart info
	cartEntity, cartItemsEntity, err := handler.cartService.GetCartAndItems(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.PaymentErr, http.StatusInternalServerError)
		return
	}

	// convert cart item to order item
	var orderItems []entity.OrderItemEntity
	for _, item := range cartItemsEntity {
		orderItems = append(orderItems, entity.OrderItemEntity{SkuId: item.SkuId, Quantity: item.Quantity})
	}

	// create order
	orderResp, err := handler.orderService.CreateNewOrder(userId, cartEntity.Total, orderItems)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("CheckoutCart: %v\n", err)
		return
	}

	// clear cart
	err = handler.cartService.ClearCartItems(cartId)
	if err != nil {
		log.Printf("CheckoutCart: order created but failed to clear cart: %s %v \n", cartId, err)
	}

	restutils.ToSuccessPayloadResponse(w, r, orderResp)
}
