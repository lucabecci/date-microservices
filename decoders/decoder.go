package decoders

import (
	"context"
	"net/http"

	"github.com/lucabecci/date-microservices/transports"
)

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transports.GetRequest
	return req, nil
}
