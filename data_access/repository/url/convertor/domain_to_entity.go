package convertor

import (
	urlEntity "github.com/likhithkp/clip/data_access/mongo/url"
	urlDomain "github.com/likhithkp/clip/domain/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DomainToEntity(domain *urlDomain.UrlDomain) (*urlEntity.UrlEntity, error) {

	var objectId, userId primitive.ObjectID
	var err error

	if domain.Id != "" {
		objectId, err = primitive.ObjectIDFromHex(domain.Id)
		if err != nil {
			return nil, err
		}
	}

	if domain.UserId != "" {
		userId, err = primitive.ObjectIDFromHex(domain.UserId)
		if err != nil {
			return nil, err
		}
	}

	return &urlEntity.UrlEntity{
		Id:        objectId,
		UserId:    userId,
		ShortUrl:  domain.ShortUrl,
		LongUrl:   domain.LongUrl,
		Title:     domain.Title,
		Code:      domain.Code,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}, nil
}
