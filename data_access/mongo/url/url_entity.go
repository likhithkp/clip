package url

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlEntity struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	ShortUrl  string             `bson:"shortUrl"`
	LongUrl   string             `bson:"longUrl"`
	Code      string             `bson:"code"`
	Title     string             `bson:"title"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
