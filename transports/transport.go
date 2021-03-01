package transports

type getRequest struct{}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json: "err,omitempty"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}
