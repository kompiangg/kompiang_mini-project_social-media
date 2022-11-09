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

func DeactivateUser(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		err := service.DeactivateUser(c.Request().Context(), userCtx.Username)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the deactivate user service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "User Deactivate!",
		})
	}
}

func EditProfile(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.EditProfileRequest
		var err error
		// err := c.Bind(&req)
		// if err != nil {
		// 	return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
		// 		Err:    errors.ErrBadRequest,
		// 		Detail: "Content type must be application/json",
		// 	})
		// }

		userCtx := authutils.UserFromRequestContext(c)

		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		req = dto.EditProfileRequest{
			Email:       c.FormValue("email"),
			DisplayName: c.FormValue("display_name"),
			Bio:         c.FormValue("bio"),
			BirthDate:   c.FormValue("birth_date"),
		}

		profilePicture, err := c.FormFile("profile_picture")
		if profilePicture != nil {
			if err != nil {
				log.Println("[HANDLER ERROR]", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		}

		var profilePictureFileName *string
		if profilePicture != nil {
			profilePictureFileName, err = httputils.HandleFileForm(profilePicture)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrInternalServer,
				})
			}
		}

		res, err := service.EditUserService(c.Request().Context(), userCtx, &req, profilePictureFileName)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the edit user service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: res,
		})
	}
}

func GetProfile(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		var resp interface{}
		var err error

		if username == userCtx.Username {
			resp, err = service.GetMyProfile(c.Request().Context(), userCtx.Username)
			if err != nil {
				log.Println("[HANDLER ERROR] While call get my profile service:", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		} else {
			resp, err = service.GetOtherUserProfile(c.Request().Context(), username)
			if err != nil {
				log.Println("[HANDLER ERROR] While call get my profile service:", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: resp,
		})
	}
}
