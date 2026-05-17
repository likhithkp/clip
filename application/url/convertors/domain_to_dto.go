package convertors

import (
	"time"

	"github.com/likhithkp/clip/application/url/dto"
	"github.com/likhithkp/clip/domain/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertUrlDtoToDomain(urlDto *dto.CreateUrlDto, userId string) (*url.UrlDomain, error) {
	var title string
	if urlDto.Title == "" {
		title = urlDto.LongUrl
	}

	return &url.UrlDomain{
		Id:        primitive.NewObjectID().Hex(),
		UserId:    userId,
		LongUrl:   urlDto.LongUrl,
		Title:     title,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
