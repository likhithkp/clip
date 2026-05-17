package url

import (
	"time"
)

type UrlDomain struct {
	Id        string
	UserId    string
	ShortUrl  string
	LongUrl   string
	Code      string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
