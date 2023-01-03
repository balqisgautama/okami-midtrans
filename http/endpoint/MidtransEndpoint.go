package endpoint

import (
	"github.com/balqisgautama/okami-midtrans/http/service"
	"net/http"
)

type midtransEndpoint struct {
	AbstractEndpoint
}

var MidtransEndpoint = midtransEndpoint{}.New()

func (input midtransEndpoint) New() (output midtransEndpoint) {
	output.FileName = "HealthEndpoint.go"
	return
}

func (input midtransEndpoint) Transaction(responseWriter http.ResponseWriter, request *http.Request) {
	input.ServeEndpoint(service.MidtransService.Transactions, responseWriter, request)
}
