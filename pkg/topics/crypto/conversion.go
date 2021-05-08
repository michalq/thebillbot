package crypto

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/michalq/thebillbot/internal/crypto"
	"github.com/michalq/thebillbot/internal/messenger"
)

type Conversion struct {
	priceProvider crypto.PriceProvider
}

func NewConversion(priceProvider crypto.PriceProvider) *Conversion {
	return &Conversion{priceProvider}
}

func (c *Conversion) Answer(ctx context.Context, message string) []messenger.Message {

	pattern := regexp.MustCompile(`([a-zA-Z]{3}) to ([a-zA-Z]{3})`)
	allIndexes := pattern.FindAllStringSubmatch(strings.ToLower(message), -1)
	answers := make([]messenger.Message, 0)
	if len(allIndexes) > 0 {
		for _, submatch := range allIndexes {
			price, err := c.priceProvider.CurrentPrice(ctx, submatch[1], submatch[2])
			if err != nil {
				answers = append(answers, messenger.Message{Content: fmt.Sprintf("%+v", err.Error())})
			} else {
				answers = append(answers, messenger.Message{Content: fmt.Sprintf("💰 1[%s] costs %s[%s]", strings.ToUpper(submatch[1]), price.Amount(), price.Currency())})
			}
		}
		return answers
	}
	return answers
}

func (c *Conversion) Promote() string {
	return `Example "LTC to USD"`
}

func (c *Conversion) Name() string {
	return "Crypto conversion"
}
