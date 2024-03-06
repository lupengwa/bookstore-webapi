package order

import (
	"bookstore-webapi/internal/api/restutils"
	"log"
	"net/http"
)

const OrderApiBasePath = "/order"
const OrderApiResourcePath = OrderApiBasePath + "/{orderId}"

// ApiHandler is the API request handler for the order domain
type ApiHandler struct {
	service Service
}

func NewApiHandler(service Service) *ApiHandler {
	if service == nil {
		log.Panic("order service can't be nil")
	}
	return &ApiHandler{service: service}
}

func (handler *ApiHandler) GetRestUriToHandlerConfig() map[restutils.RestApiUriKey]http.HandlerFunc {
	return map[restutils.RestApiUriKey]http.HandlerFunc{
		restutils.RestApiUriKey{
			HttpMethod: http.MethodGet,
			Path:       OrderApiResourcePath,
		}: handler.ListOrderItems,
	}
}

// todo
func (handler *ApiHandler) ListOrderItems(w http.ResponseWriter, r *http.Request) {

}
