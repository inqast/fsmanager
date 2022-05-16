package models

import "database/sql"

type User struct {
	ID        int
	Name      string
	Pwd       string
	CreatedAt sql.NullTime
}

type Subscription struct {
	ID                int
	OwnerID           int
	ServiceName       string
	Capacity          int
	PriceInCentiUnits int
	PaymentDate       sql.NullTime
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
