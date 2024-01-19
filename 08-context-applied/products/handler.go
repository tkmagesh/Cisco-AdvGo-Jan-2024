package products

import (
	"encoding/json"

	"net/http"

	log "github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"req-id": r.Context().Value("request-id"),
	}).Info("[Products Handler]")

	ps := NewProductsService()
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(ps.GetAll(r.Context())); err != nil {
			http.Error(w, "error encoding data", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
		}
		ps.AddNew(r.Context(), newProduct)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
