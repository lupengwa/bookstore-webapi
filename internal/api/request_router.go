package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const BasePath = "/api"

// RequestRouteConfigurer has the logic to configure the request routing in the app
// and uses the injected factory to provide the API handler for the specific path
type RequestRouteConfigurer struct {
	factory HandlerFactory
}

func NewRequestRouteConfigurer(factory HandlerFactory) *RequestRouteConfigurer {
	if factory == nil {
		log.Panic("HandlerFactory can't be nil")
	}
	return &RequestRouteConfigurer{factory: factory}
}

func (router *RequestRouteConfigurer) Configure(route chi.Router) (err error) {

	route.Route(BasePath, func(r chi.Router) {

		// Set up all path and API handlers mapping provided by the factory
		for key, handler := range router.factory.GetApiUriToHandler() {
			switch key.HttpMethod {
			case http.MethodGet:
				r.Get(key.Path, handler)
			case http.MethodPost:
				r.Post(key.Path, handler)
			case http.MethodPut:
				r.Put(key.Path, handler)
			case http.MethodDelete:
				r.Delete(key.Path, handler)
			case http.MethodPatch:
				r.Patch(key.Path, handler)
			default:
				err = fmt.Errorf("unsupported HTTP method: %s", key.HttpMethod)
			}
		}
	})
	return
}
