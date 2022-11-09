package middleware

import (
	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func AdminSecretTokenAuth(config *config.Admin) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Secret-Token")
			if authorization == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    errors.ErrBadRequest,
					Detail: []string{"Secret-Token header value couldn't be empty"},
				})
			}

			if authorization != config.SecretToken {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrUnauthorized,
				})
			}

			return next(c)
		}
	}
}
