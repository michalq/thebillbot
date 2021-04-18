package crypto

import "context"

type PriceProvider interface {
	CurrentPrice(ctx context.Context, product, currency string) (*Price, error)
}
