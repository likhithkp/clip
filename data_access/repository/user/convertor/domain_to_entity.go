package convertor

import (
	userEntity "github.com/likhithkp/clip/data_access/mongo/user"
	userDomain "github.com/likhithkp/clip/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DomainToEntity(domain *userDomain.UserDomain) (*userEntity.UserEntity, error) {

	var objectId primitive.ObjectID
	var err error

	if domain.Id != "" {
		objectId, err = primitive.ObjectIDFromHex(domain.Id)
		if err != nil {
			return nil, err
		}
	}

	return &userEntity.UserEntity{
		Id:        objectId,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Password:  domain.Password,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}, nil
}
