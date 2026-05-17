package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	urlRepository "github.com/likhithkp/clip/data_access/repository/url"
	"github.com/likhithkp/clip/utils/other"
)

type GetUrlHandler struct {
	utils         *other.ResponseStruct
	urlRepository *urlRepository.UrlRepository
}

func NewGetUrlHandler(
	utils *other.ResponseStruct,
	urlRepository *urlRepository.UrlRepository,
) *GetUrlHandler {
	return &GetUrlHandler{
		utils:         utils,
		urlRepository: urlRepository,
	}
}

func (handler *GetUrlHandler) GetUrl(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return handler.utils.Response(ctx, http.StatusBadRequest, false, "Code not provided", nil)
	}

	urlDomain, err := handler.urlRepository.GetUrlByCode(ctx.Context(), code)
	if err != nil {
		log.Printf("error[GetUrl]: %v", err.Error())
		return handler.utils.Response(ctx, http.StatusInternalServerError, false, "Internal server error", nil)
	}

	return ctx.Redirect(urlDomain.LongUrl, http.StatusMovedPermanently)
}
