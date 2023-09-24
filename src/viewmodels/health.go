package viewmodels

const (
	statusOk = "OK"
)

type healthResponse struct {
	Status string `json:"status"`
}

func NewHealthOk() *healthResponse {
	return &healthResponse{
		Status: statusOk,
	}
}
