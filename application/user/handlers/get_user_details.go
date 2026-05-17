package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/user/convertors"
	userRepository "github.com/likhithkp/clip/data_access/repository/user"
	"github.com/likhithkp/clip/utils/other"
)

type GetUserDetailsHandler struct {
	utils          *other.ResponseStruct
	userRepository *userRepository.UserRepository
}

func NewGetUserDetailsHandler(
	utils *other.ResponseStruct,
	userRepository *userRepository.UserRepository,
) *GetUserDetailsHandler {
	return &GetUserDetailsHandler{
		utils:          utils,
		userRepository: userRepository,
	}
}

func (handler *GetUserDetailsHandler) GetUserDetails(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userID")
	if userId == nil {
		log.Println("error[GetUserDetails]", "Error while getting userID from middleware")
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	userIdStr := userId.(string)

	userDomain, err := handler.userRepository.GetUserById(ctx.Context(), userIdStr)
	if err != nil {
		log.Printf("error[GetUserDetails]: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}
	if userDomain == nil {
		return handler.utils.Response(ctx, http.StatusNotFound, false, "User doesn't exist", nil)
	}

	userDetailsDto := convertors.ConvertDomainToDetailsDto(userDomain)
	return handler.utils.Response(ctx, http.StatusOK, true, "User details fetched successfully", userDetailsDto)
}
