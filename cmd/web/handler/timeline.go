package handler

import (
	"log"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func GetTimeline(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		timeline, err := service.GetTimeline(c.Request().Context(), userCtx)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the get timeline service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: timeline,
		})
	}
}
