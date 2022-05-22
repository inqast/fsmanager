package message

import "fmt"

type SubscriptionsResponse struct {
	UserName      string
	Subscriptions []*Subscription
}

func (s *SubscriptionsResponse) String() string {
	result := fmt.Sprintf("Here all the subscriptions for %s: \n", s.UserName)

	for _, subscription := range s.Subscriptions {
		result += subscription.String()
	}

	return result
}
