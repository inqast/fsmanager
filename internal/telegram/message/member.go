package message

import "fmt"

type Member struct {
	UserID  int
	Name    string
	IsOwner bool
	IsPaid  bool
}

func (m *Member) String() string {
	result := fmt.Sprintf("@%s isPaid: %t\n",
		m.Name,
		m.IsPaid,
	)

	return result
}
