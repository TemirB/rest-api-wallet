package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/TemirB/rest-api-wallet/internal/models"
	"github.com/TemirB/rest-api-wallet/internal/service"
	"github.com/google/uuid"
)

func HandleWallet(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	if r.Method == http.MethodPost {
		var req models.WalletRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = svc.UpdateWalletBalance(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Operation successful"))
	}
}

func GetWallet(w http.ResponseWriter, r *http.Request, svc *service.Service) {
	path := strings.TrimPrefix(r.URL.Path, "/wallet/")
	if path == "" {
		http.Error(w, "Wallet ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(path)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	balance, err := svc.GetWalletBalance(id)
	if err != nil {
		http.Error(w, "Error retrieving wallet balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"balance": balance}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
