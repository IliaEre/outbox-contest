package handler

import (
	"encoding/json"
	"net/http"
	"order-service/domain"
)

func (h *OrderHandler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	userUUID := r.Header.Get("USER_UUID")
	if userUUID == "" {
		http.Error(w, "USER_UUID is important", http.StatusInternalServerError)
		return
	}

	var orderRequest domain.Order

	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID, err := h.service.Save(userUUID, &orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"order_id": orderID})
}
