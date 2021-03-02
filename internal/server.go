package internal

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/lucabecci/date-microservices/endpoints"
	"github.com/lucabecci/date-microservices/internal/middlewares"
	"github.com/lucabecci/date-microservices/utils"
)

func NewServer(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middlewares.CommonMiddleware)

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		utils.DecodeStatusRequest,
		utils.EncodeResponse,
	))

	r.Methods("GET").Path("/get").Handler(httptransport.NewServer(
		endpoints.GetEndpoint,
		utils.DecodeGetRequest,
		utils.EncodeResponse,
	))

	r.Methods("POST").Path("/validate").Handler(httptransport.NewServer(
		endpoints.ValidateEndpoint,
		utils.DecodeValidateRequest,
		utils.EncodeResponse,
	))

	return r
}
