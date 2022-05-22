package message

import "fmt"

type Success struct {
	Id int
}

func (s *Success) String() string {
	result := "Success"
	if s.Id != 0 {
		result += fmt.Sprintf(" id: %d", s.Id)
	}
	return result
}
