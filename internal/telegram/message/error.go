package message

type Error struct {
	Msg string
}

func (e *Error) String() string {
	return e.Msg
}
