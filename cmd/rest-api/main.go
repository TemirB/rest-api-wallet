package main

import (
	"log"
	"net/http"

	"github.com/TemirB/rest-api-wallet/internal/handler"
	"github.com/TemirB/rest-api-wallet/internal/repository"
	"github.com/TemirB/rest-api-wallet/internal/service"
	"github.com/TemirB/rest-api-wallet/pkg/db"
)

func main() {
	datb := db.Connect()

	repo := &repository.Repository{
		DB: datb,
	}

	svc := &service.Service{
		Repo: repo,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/wallet", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleWallet(w, r, svc)
	})
	mux.HandleFunc("/wallets/", func(w http.ResponseWriter, r *http.Request) {
		handler.GetWallet(w, r, svc)
	})

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
