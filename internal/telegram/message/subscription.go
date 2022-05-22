package message

import "fmt"

type Subscription struct {
	Id         int
	Service    string
	Owner      string
	Cost       float64
	PaymentDay int
	Members    []*Member
	Capacity   int
	Share      float64
	IsPaid     bool
}

func (s *Subscription) String() string {
	result := fmt.Sprintf(`
		id: %d service: %s
		owner: %s
		cost: %.2f
		paymentDay: %d
		capacity: %d/%d
		share: %.2f
		isPaid: %t
		members:
	`,
		s.Id, s.Service,
		s.Owner,
		s.Cost,
		s.PaymentDay,
		len(s.Members), s.Capacity,
		s.Share,
		s.IsPaid,
	)
	for _, member := range s.Members {
		result += member.String()
	}

	return result
}
