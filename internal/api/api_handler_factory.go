package api

import (
	"bookstore-webapi/internal/api/cart"
	"bookstore-webapi/internal/api/restutils"
	"bookstore-webapi/internal/platform/db"
	"fmt"
	"log"
	"net/http"
)

type HandlerFactory interface {
	GetApiUriToHandler() map[restutils.RestApiUriKey]http.HandlerFunc
}

// HandlerFactoryImpl is the implementation of the simple factory pattern.
// - Provides one single method to retrieve the mapping of URI and the associated API handler function
// - Code to instantiate the API handler objects and their dependent objects
type HandlerFactoryImpl struct {
	dataSource          *db.DataSource
	apiUriToHandlerFunc map[restutils.RestApiUriKey]http.HandlerFunc
}

func NewHandlerFactory(dataSource *db.DataSource) *HandlerFactoryImpl {
	factory := &HandlerFactoryImpl{
		dataSource:          dataSource,
		apiUriToHandlerFunc: make(map[restutils.RestApiUriKey]http.HandlerFunc),
	}

	err := factory.initCartHandler()
	if err != nil {
		log.Panic("Failed to load cart handler")
	}

	err = factory.initOrderHandler()
	if err != nil {
		log.Panic("Failed to load order handler")
	}

	err = factory.initProductHandler()
	if err != nil {
		log.Panic("Failed to load product handler")
	}

	return factory
}

func (factory *HandlerFactoryImpl) GetApiUriToHandler() map[restutils.RestApiUriKey]http.HandlerFunc {
	return factory.apiUriToHandlerFunc
}

func (factory *HandlerFactoryImpl) registerUriPathForHandler(handler Handler) error {
	if len(factory.apiUriToHandlerFunc) == 0 {
		factory.apiUriToHandlerFunc = make(map[restutils.RestApiUriKey]http.HandlerFunc)
	}
	for key, handlerFunc := range handler.GetRestUriToHandlerConfig() {
		if _, found := factory.apiUriToHandlerFunc[key]; found {
			return fmt.Errorf("API endbpoint[%v: %v]is already registered at ApiHandlerFactory", key.HttpMethod, key.Path)
		}
		factory.apiUriToHandlerFunc[key] = handlerFunc
	}
	return nil
}

func (factory *HandlerFactoryImpl) initCartHandler() error {
	cartRepo := cart.NewRepository(factory.dataSource)
	cartService := cart.NewService(cartRepo)
	cartHandler := cart.NewApiHandler(cartService)
	err := factory.registerUriPathForHandler(cartHandler)
	if err != nil {
		log.Println("failed to register cart handler")
		return err
	}
	return nil
}

func (factory *HandlerFactoryImpl) initOrderHandler() error {
	return nil
}

func (factory *HandlerFactoryImpl) initProductHandler() error {
	return nil
}
