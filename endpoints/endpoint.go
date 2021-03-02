package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucabecci/date-microservices/services"
	"github.com/lucabecci/date-microservices/transports"
)

type Endpoints struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

//makes
func MakeGetEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transports.GetRequest)
		d, err := srv.Get(ctx)
		if err != nil {
			return transports.GetResponse{Date: d, Err: err.Error()}, nil
		}
		return transports.GetResponse{Date: d, Err: ""}, nil
	}
}

func MakeStatusEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transports.StatusRequest) // we really just need the request, we don't use any value from it
		s, err := srv.Status(ctx)
		if err != nil {
			return transports.StatusResponse{Status: s}, err
		}
		return transports.StatusResponse{Status: s}, nil
	}
}

func MakeValidateEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.ValidateRequest)
		b, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return transports.ValidateResponse{Valid: b, Err: err.Error()}, nil
		}
		return transports.ValidateResponse{Valid: b, Err: ""}, nil
	}
}

//Get endpoint
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := transports.GetRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transports.GetResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

// Status endpoint
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := transports.StatusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(transports.StatusResponse)
	return statusResp.Status, nil
}

// Validate endpoint
func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := transports.ValidateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(transports.ValidateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
