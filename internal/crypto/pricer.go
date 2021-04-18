package crypto

type Price struct {
	currency string
	amount   string
}

func NewPrice(currency, amount string) *Price {
	return &Price{currency, amount}
}

func (p *Price) Amount() string {
	return p.amount
}

func (p *Price) Currency() string {
	return p.currency
}
