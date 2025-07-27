package main

import (
	"fmt"
	"ms-go-notification/internal/domain/campaign"
	"ms-go-notification/internal/endpoints"
	"ms-go-notification/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
