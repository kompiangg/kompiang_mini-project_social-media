package handler

import (
	"log"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func GetFollowers(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		followers, err := service.GetFollowers(c.Request().Context(), username)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get followers service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *followers,
		})
	}
}

func GetFollowing(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		followings, err := service.GetFollowings(c.Request().Context(), username)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get followers service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: *followings,
		})
	}
}

func FollowOtherUser(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.FollowRequest
		err := c.Bind(&req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"Content type must be application/json"},
			})
		}

		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		err = service.FollowOtherUser(c.Request().Context(), &req, userCtx)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get child comment service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Followed",
		})
	}
}
