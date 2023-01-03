package service

import (
	"bytes"
	"encoding/json"
	"github.com/balqisgautama/okami-midtrans/config"
	"github.com/balqisgautama/okami-midtrans/constanta"
	"github.com/balqisgautama/okami-midtrans/dto"
	"github.com/balqisgautama/okami-midtrans/dto/res"
	"github.com/balqisgautama/okami-midtrans/model"
	"github.com/balqisgautama/okami-midtrans/util"
	"io/ioutil"
	"net/http"
)

type AbstractService struct {
	FileName string
	FuncName string
}

func (input AbstractService) ReadBody(request *http.Request) (stringBody string, output res.APIResponse) {
	var errorS error
	input.FuncName = "ReadBody"
	input.FileName = "AbstractService.go"

	if request.Method != "GET" {
		stringBody, _, errorS = util.ReadBody(request)
		if errorS != nil {
			output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
			return
		}
	}

	return
}

func (input AbstractService) MidtransTransaction(data any) (transactionResponse res.TransactionResponse, output res.APIResponse) {
	input.FuncName = "MidtransTransaction"
	input.FileName = "AbstractService.go"

	postBody, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postBody)
	url := constanta.MIDTRANS_URL_TRANSACTION

	client := http.Client{}
	req, err := http.NewRequest(constanta.RequestPOST, url, responseBody)
	if err != nil {
		output = model.GenerateMidtransServerError(input.FileName, input.FuncName, err)
		return
	}

	token := util.Base64Encrypt(config.ApplicationConfiguration.GetMidtransServerKey() + ":")

	req.Header = http.Header{
		constanta.HeaderKeyContentType:    {constanta.HeaderValueContentTypeJSON},
		constanta.AuthorizationHeaderName: {"Basic " + token},
	}

	resp, err := client.Do(req)
	if err != nil {
		output = model.GenerateMidtransServerError(input.FileName, input.FuncName, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		output = model.GenerateMidtransServerError(input.FileName, input.FuncName, err)
		return
	}
	sb := string(body)

	err = json.Unmarshal([]byte(sb), &transactionResponse)
	if err != nil {
		output = model.GenerateMidtransServerError(input.FileName, input.FuncName, err)
		return
	}

	return
}
