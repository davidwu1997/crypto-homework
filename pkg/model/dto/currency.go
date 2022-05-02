package dto

type CreateCurrencyReq struct {
	Code string `json:"currency_code"`
	Name string `json:"currency_name"`
}
