package routes

import (
	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest"
	"github.com/go-chi/chi/v5"
)

func RegisterOrderRoutes(r chi.Router, handler *rest.Handler) {
	r.Route("/v1/order", func(r chi.Router) {
		r.Post("/", handler.CreateOrder)
		r.Get("/", handler.ListOrders)
	})
}
