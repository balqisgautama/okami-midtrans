package service

import (
	"encoding/json"
	"github.com/balqisgautama/okami-midtrans/dto"
	"github.com/balqisgautama/okami-midtrans/dto/req"
	"github.com/balqisgautama/okami-midtrans/dto/res"
	"net/http"
)

type midtransService struct {
	AbstractService
}

var MidtransService = midtransService{}.New()

func (input midtransService) New() (output midtransService) {
	output.FileName = "HealthService.go"
	return
}

func (input midtransService) Transactions(request *http.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "Transactions"

	result, output_ := input.readBodyAndValidateTransaction(request, req.ValidateTransaction)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	createTransaction, output_ := input.MidtransTransaction(result)
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Data = createTransaction
	return
}

func (input midtransService) readBodyAndValidateTransaction(request *http.Request, validation func(input *req.GenerateTransactionRequest) (output res.APIResponse)) (inputStruct req.GenerateTransactionRequest, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateTransaction"
	var stringBody string

	stringBody, output = input.ReadBody(request)
	if output.Status.Code != "" {
		return
	}

	if stringBody != "" {
		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
		if errorS != nil {
			output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
			return
		}
	}
	output = validation(&inputStruct)
	if output.Status.Code != "" {
		return
	}

	return
}
