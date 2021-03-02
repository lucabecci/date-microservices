package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lucabecci/date-microservices/transports"
)

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transports.GetRequest
	return req, nil
}

func DecodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transports.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transports.StatusRequest
	return req, nil
}
