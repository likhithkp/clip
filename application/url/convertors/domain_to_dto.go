package convertors

import (
	"fmt"
	"time"

	"github.com/likhithkp/clip/application/url/dto"
	"github.com/likhithkp/clip/domain/url"
	"github.com/likhithkp/clip/utils/other"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertUrlDtoToDomain(urlDto *dto.CreateUrlDto, userId string) (*url.UrlDomain, error) {

	code, err := other.GenerateShortCode(urlDto.Code)
	if err != nil {
		return nil, err
	}

	shortUrl := fmt.Sprintf("http://localhost:8080/api/v1/urls/%s", code)

	var title string
	if urlDto.Title == "" {
		title = urlDto.LongUrl
	}

	return &url.UrlDomain{
		Id:        primitive.NewObjectID().Hex(),
		UserId:    userId,
		LongUrl:   urlDto.LongUrl,
		ShortUrl:  shortUrl,
		Title:     title,
		Code:      code,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
