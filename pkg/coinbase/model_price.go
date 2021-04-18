package coinbase

type ModelPricesResponse struct {
	Data ModelPrice `json:"data"`
}

type ModelPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
