package handler

import (
	"log"
	"net/http"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func AdminCreateGeneralNotifications(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.AdminNotificationRequest
		err := c.Bind(&req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"Content type must be application/json"},
			})
		}

		if req.Content == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"Content shouldnt be empty"},
			})
		}

		err = service.CreateGeneralNotifications(c.Request().Context(), &req)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the create general notifications service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: http.StatusCreated,
			Data: "General Notification Created!",
		})
	}
}

func GetNotifications(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		notifications, err := service.GetNotificationsByUsername(c.Request().Context(), userCtx.Username)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: notifications,
		})
	}
}
