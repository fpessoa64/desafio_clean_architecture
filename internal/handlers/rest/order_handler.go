package rest

import (
	"encoding/json"
	"net/http"

	"github.com/fpessoa64/desafio_clean_arch/internal/entities"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
)

type Handler struct {
	UC *usecase.OrderUsecase
}

func NewHandler(uc *usecase.OrderUsecase) *Handler {
	return &Handler{UC: uc}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var in entities.Order
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := h.UC.Create(r.Context(), &in); err != nil {
		http.Error(w, "error creating order", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(in)
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.UC.List(r.Context())
	if err != nil {
		http.Error(w, "error listing orders", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
