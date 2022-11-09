package middleware

import (
	"strings"

	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/tokenutils"
	"github.com/labstack/echo/v4"
)

func MustAuthorized(config *config.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    errors.ErrBadRequest,
					Detail: []string{"Authorization header value couldn't be empty"},
				})
			}

			splitted := strings.SplitAfter(authorization, "Bearer ")
			if len(splitted) != 2 {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    errors.ErrBadRequest,
					Detail: []string{"Bearer format is not valid"},
				})
			}

			if splitted[1] == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    errors.ErrBadRequest,
					Detail: []string{"Bearer value is couldn't empty"},
				})
			}

			accessToken := splitted[1]

			user, err := tokenutils.ValidateAccessToken(config, accessToken)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrUnauthorized,
				})
			}

			c.Set(dto.AccountCTXKey, user)

			return next(c)
		}
	}
}
