package res

type TransactionResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}
