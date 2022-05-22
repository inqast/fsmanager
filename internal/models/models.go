package models

import "database/sql"

type User struct {
	ID         int
	Name       string
	Pwd        string
	TelegramID int
	CreatedAt  sql.NullTime
}

type Subscription struct {
	ID                int
	ChatID            int
	ServiceName       string
	Capacity          int
	PriceInCentiUnits int
	PaymentDay        int
	CreatedAt         sql.NullTime
}

type Subscriber struct {
	ID             int
	UserID         int
	SubscriptionID int
	IsPaid         bool
	IsOwner        bool
	CreatedAt      sql.NullTime
}
