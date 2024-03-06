package application

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (b *BookStoreApp) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/orders", loadOrderRoutes)
	b.router = router
}

func loadOrderRoutes(router chi.Router) {
	//orderHandler := &handler.Order{}
	//router.Post("/", orderHandler.Create)
	//router.Get("/", orderHandler.List)
	//router.Get("/{id}", orderHandler.GetByID)
	//router.Put("/{id}", orderHandler.UpdateByID)
	//router.Delete("/{id}", orderHandler.DeleteByID)
}
