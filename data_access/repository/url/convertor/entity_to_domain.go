package convertor

import (
	urlEntity "github.com/likhithkp/clip/data_access/mongo/url"
	urlDomain "github.com/likhithkp/clip/domain/url"
)

func EntityToDomain(entity *urlEntity.UrlEntity) *urlDomain.UrlDomain {
	return &urlDomain.UrlDomain{
		Id:        entity.Id.Hex(),
		UserId:    entity.UserId.Hex(),
		ShortUrl:  entity.ShortUrl,
		LongUrl:   entity.LongUrl,
		Code:      entity.Code,
		Title:     entity.Title,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
