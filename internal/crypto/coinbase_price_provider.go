package crypto

import (
	"context"
	"fmt"
	"strings"

	"github.com/michalq/thebillbot/pkg/coinbase"
)

type CoinbasePriceProvider struct {
	coinbaseClient *coinbase.Client
}

func NewCoinbasePriceProvider(coinbaseClient *coinbase.Client) *CoinbasePriceProvider {
	return &CoinbasePriceProvider{coinbaseClient}
}

func (c *CoinbasePriceProvider) CurrentPrice(ctx context.Context, product, currency string) (*Price, error) {
	resp, err := c.coinbaseClient.PricesSpot(ctx, fmt.Sprintf("%s-%s", strings.ToUpper(product), strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}

	return NewPrice(resp.Data.Currency, resp.Data.Amount), nil
}
