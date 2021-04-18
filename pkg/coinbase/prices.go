package coinbase

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) PricesBuy(ctx context.Context, pair string) (*ModelPricesResponse, error) {
	return c.prices(ctx, pair, "buy")
}

func (c *Client) PricesSell(ctx context.Context, pair string) (*ModelPricesResponse, error) {
	return c.prices(ctx, pair, "sell")
}

func (c *Client) PricesSpot(ctx context.Context, pair string) (*ModelPricesResponse, error) {
	return c.prices(ctx, pair, "spot")
}

func (c *Client) prices(ctx context.Context, pair, priceType string) (*ModelPricesResponse, error) {

	response := &ModelPricesResponse{}
	if err := c.request(ctx, http.MethodGet, fmt.Sprintf("/v2/prices/%s/%s", pair, priceType), response); err != nil {
		return nil, err
	}

	return response, nil
}
