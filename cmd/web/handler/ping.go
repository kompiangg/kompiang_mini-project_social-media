package handler

import (
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func PingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "pong",
		})
	}
}
