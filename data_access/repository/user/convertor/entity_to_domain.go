package convertor

import (
	userEntity "github.com/likhithkp/clip/data_access/mongo/user"
	userDomain "github.com/likhithkp/clip/domain/user"
)

func EntityToDomain(entity *userEntity.UserEntity) *userDomain.UserDomain {
	return &userDomain.UserDomain{
		Id:        entity.Id.Hex(),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Password:  entity.Password,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
