package application

import (
	"bookstore-webapi/internal/api"
	"bookstore-webapi/internal/platform/db"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

// BookStoreApp is the main application
type BookStoreApp struct {
	router http.Handler
	config BookStoreApiServiceProperty
}

func NewBookStoreApp(config BookStoreApiServiceProperty) *BookStoreApp {

	// Init DB connection
	dataSource := db.NewDataSource(config.PgDbConfig)
	handlerFactory := api.NewHandlerFactory(dataSource)

	// Map request to corresponding handler
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	api.NewRequestRouteConfigurer(handlerFactory).Configure(router)

	bookStoreApp := &BookStoreApp{
		config: config,
		router: router,
	}
	return bookStoreApp
}

func (b *BookStoreApp) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", b.config.ServerPort),
		Handler: b.router,
	}

	var err error

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
	return nil
}
