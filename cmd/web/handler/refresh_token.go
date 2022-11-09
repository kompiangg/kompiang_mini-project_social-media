package handler

import (
	"log"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func RefreshToken(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var refreshToken dto.RefreshTokenRequest
		err := c.Bind(&refreshToken)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: "Content type must be application/json",
			})
		}

		if refreshToken.RefreshToken == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: "Refresh Token field couldn't be empty",
			})
		}

		accessToken, err := service.RefreshTokenService(c.Request().Context(), &refreshToken)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the refresh token service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: dto.UserSuccessLoginResponse{
				AccessToken: accessToken.AccessToken,
			},
		})
	}
}
