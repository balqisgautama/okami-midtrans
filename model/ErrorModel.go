package model

import (
	"github.com/balqisgautama/okami-midtrans/config"
	"github.com/balqisgautama/okami-midtrans/dto/res"
	"strings"
)

var resourceID string

func GenerateMidtransServerError(fileName string, funcName string, causedBy error) (output res.APIResponse) {
	resourceID = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID())
	output.Status.Success = false
	output.Status.Code = resourceID + "-370001-MIDTRANS-SERVER"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName}
	return
}
