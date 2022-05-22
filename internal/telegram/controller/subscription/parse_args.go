package subscription

import (
	"errors"
	"strconv"
	"strings"

	"github.com/inqast/fsmanager/internal/models"
)

func (c *Controller) parseArgs(args []string) (*models.Subscription, error) {
	subscription := models.Subscription{}
	for _, arg := range args {
		kv := strings.Split(arg, "=")

		switch kv[0] {
		case "name":
			subscription.ServiceName = kv[1]
		case "cost":
			value, err := strconv.Atoi(kv[1])
			if err != nil {
				return nil, err
			}
			if value <= 0 {
				return nil, errors.New("cost should be positive number")
			}
			subscription.PriceInCentiUnits = value * 100
		case "cap":
			value, err := strconv.Atoi(kv[1])
			if err != nil {
				return nil, err
			}
			if value <= 0 {
				return nil, errors.New("cap should be positive number")
			}
			subscription.Capacity = value
		case "payday":
			value, err := strconv.Atoi(kv[1])
			if err != nil {
				return nil, err
			}
			if value <= 0 || value > 31 {
				return nil, errors.New("payday should be in range 1-31)")
			}
			subscription.PaymentDay = value
		}
	}
	return &subscription, nil
}
