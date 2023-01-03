package service

import (
	"github.com/balqisgautama/okami-midtrans/dto/res"
	"net/http"
)

type healthService struct {
	AbstractService
}

var HealthService = healthService{}.New()

func (input healthService) New() (output healthService) {
	output.FileName = "HealthService.go"
	return
}

func (input healthService) CheckingHealth(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "CheckingHealth"

	output.Status.Success = true
	output.Status.Message = "UP"
	return
}
