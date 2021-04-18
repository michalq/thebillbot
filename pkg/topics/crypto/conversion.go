package crypto

import (
	"context"
	"fmt"
	"regexp"

	"github.com/michalq/thebillbot/internal/crypto"
)

type Conversion struct {
	priceProvider crypto.PriceProvider
}

func NewConversion(priceProvider crypto.PriceProvider) *Conversion {
	return &Conversion{priceProvider}
}

func (c *Conversion) Answer(message string) []string {

	pattern := regexp.MustCompile(`([a-zA-Z]{3}) to ([a-zA-Z]{3})`)
	allIndexes := pattern.FindAllStringSubmatch(message, -1)
	answers := make([]string, 0)
	if len(allIndexes) > 0 {
		for _, submatch := range allIndexes {
			price, err := c.priceProvider.CurrentPrice(context.TODO(), submatch[1], submatch[2])
			if err != nil {
				answers = append(answers, fmt.Sprintf("%+v", err.Error()))
			} else {
				answers = append(answers, fmt.Sprintf("test %s %s", price.Amount(), price.Currency()))
			}
		}
		return answers
	}
	return answers
}
