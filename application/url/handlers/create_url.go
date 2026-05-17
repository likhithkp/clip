package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/url/convertors"
	"github.com/likhithkp/clip/application/url/dto"
	urlRepository "github.com/likhithkp/clip/data_access/repository/url"
	"github.com/likhithkp/clip/utils/other"
)

type CreateUrlHandler struct {
	utils         *other.ResponseStruct
	urlRepository *urlRepository.UrlRepository
}

func NewCreateUrlHanler(
	utils *other.ResponseStruct,
	urlRepository *urlRepository.UrlRepository,
) *CreateUrlHandler {
	return &CreateUrlHandler{
		utils:         utils,
		urlRepository: urlRepository,
	}
}

func (handler *CreateUrlHandler) CreateUrl(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userID")
	if userId == "" {
		log.Println("error[CreateUrl]", "Unable to fetch user id from ctx locals")
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	userIdStr := userId.(string)

	newUrl := new(dto.CreateUrlDto)
	err := ctx.BodyParser(newUrl)
	if err != nil {
		return handler.utils.Response(ctx, http.StatusUnprocessableEntity, false, "Error while parsing JSON body", nil)
	}
	if newUrl.LongUrl == "" {
		return handler.utils.Response(ctx, http.StatusBadRequest, false, "Missing title or link", nil)
	}

	urlDomain, err := convertors.ConvertUrlDtoToDomain(newUrl, userIdStr)
	if err != nil {
		log.Printf("error[CreateUrl]: %s", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	//TODO:Save the url using HSET in redis, make sure the code is not duplicated

	err = handler.urlRepository.UpsertUrl(ctx.Context(), urlDomain)
	if err != nil {
		log.Printf("error[CreateUrl]: %s", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	return handler.utils.Response(ctx, http.StatusCreated, true, "Short URL created successfully", nil)
}
