package command

import (
	"fmt"
	"strings"
)

type Query struct {
	Domain  string
	Command string
	Args    []string
}

func Parse(commandText string) (*Query, error) {
	commandParts := strings.Split(commandText, " ")
	query := Query{}
	query.Domain = commandParts[0]
	if len(commandParts) > 1 {
		query.Command = commandParts[1]
		query.Args = commandParts[2:]
	}

	return &query, nil
}

func (q Query) String() string {
	return fmt.Sprintf("/%s %s", q.Domain, q.Command)
}
