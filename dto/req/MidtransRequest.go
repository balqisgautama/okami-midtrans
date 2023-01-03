package req

import (
	"github.com/balqisgautama/okami-midtrans/dto"
	"github.com/balqisgautama/okami-midtrans/dto/res"
	"github.com/go-playground/validator/v10"
)

type GenerateTransactionRequest struct {
	TransactionDetails struct {
		OrderID     string `json:"order_id" validate:"required"`
		GrossAmount int64  `json:"gross_amount" validate:"required,gt=0,number"`
	} `json:"transaction_details" validate:"required"`
	CustomerDetails struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,min=5,email"`
		Phone     string `json:"phone" validate:"omitempty"`
	} `json:"customer_details" validate:"omitempty"`
}

func ValidateTransaction(inputStruct *GenerateTransactionRequest) (output res.APIResponse) {
	funcName = "ValidateTransaction"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output = dto.GenerateValidationFailed(err, filename, funcName)
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output = dto.GenerateValidationFailed(errV, filename, funcName)
			return
		}
		return
	}
	output.Status.Success = true
	return
}
