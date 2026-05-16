package user

import "time"

type UserDomain struct {
	Id        string
	FirstName string
	LastName  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
