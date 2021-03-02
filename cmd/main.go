package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lucabecci/date-microservices/endpoints"
	"github.com/lucabecci/date-microservices/internal"
	"github.com/lucabecci/date-microservices/services"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()

	srv := services.NewService()

	endpoints := endpoints.Endpoints{
		GetEndpoint:      endpoints.MakeGetEndpoint(srv),
		StatusEndpoint:   endpoints.MakeStatusEndpoint(srv),
		ValidateEndpoint: endpoints.MakeValidateEndpoint(srv),
	}

	handler := internal.NewServer(ctx, endpoints)
	logger.Log("msg", "HTTP", "addr", ":4000")
	logger.Log("err", http.ListenAndServe(":4000", handler))

}
