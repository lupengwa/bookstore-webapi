package cart

import (
	"bookstore-webapi/internal/api/cart/dto"
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
	service Service
}

func NewApiHandler(cartService Service) *ApiHandler {
	if cartService == nil {
		log.Panic("Cart repo can't be nil")
	}
	return &ApiHandler{service: cartService}
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

	cart, err := handler.service.GetCartItems(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("ListCartItems: %v\n", err)
		return
	}
	restutils.ToSuccessPayloadResponse(w, r, cart)
}

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

	err = handler.service.ClearCart(cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("UpdateCart: %s failed cause: %s", cartId, err)
		return
	}

	updatedCart, err := handler.service.AddCartItems(cartDto, cartId)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.ServerErr, http.StatusInternalServerError)
		log.Printf("UpdateCart: failed to update cart item. cause: %v\n", err)
		return
	}
	restutils.ToSuccessPayloadResponse(w, r, updatedCart)

}

func (handler *ApiHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	userId, err := restutils.ValidateUser(r)
	if err != nil {
		restutils.ToErrorResponse(w, r, apperror.InvalidUserErr, http.StatusUnauthorized)
		return
	}
	// check if cart already exists
	cartId := userId

	// use Location to inform client the resource directory
	w.Header().Set("Location", fmt.Sprintf("/cart/%s", cartId))

	isNewCart, cartDto, err := handler.service.CreateNewCart(cartId)
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

func (handler *ApiHandler) CheckoutCart(w http.ResponseWriter, r *http.Request) {

}
