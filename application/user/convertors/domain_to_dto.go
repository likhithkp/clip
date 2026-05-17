package convertors

import (
	"github.com/likhithkp/clip/application/user/dto"
	"github.com/likhithkp/clip/domain/user"
)

func ConvertDomainToDetailsDto(userDomain *user.UserDomain) *dto.UserDetailsDto {
	return &dto.UserDetailsDto{
		Id:        userDomain.Id,
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Email:     userDomain.Email,
	}
}
