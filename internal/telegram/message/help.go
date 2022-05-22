package message

type Help struct {
	Message string
}

func (h Help) String() string {
	return h.Message
}
