package handler

import (
	"log"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func Login(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var registerReq dto.UserLoginRequest
		if err := c.Bind(&registerReq); err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: "Content type must be application/json",
			})
		}

		detail, err := validateLogin(&registerReq)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    err,
				Detail: detail,
			})
		}

		token, err := service.LoginService(c.Request().Context(), &registerReq)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the login service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: token,
		})
	}
}

func validateLogin(req *dto.UserLoginRequest) ([]string, error) {
	var badRequestDetails []string

	if req.Username == "" {
		badRequestDetails = append(badRequestDetails, "Username cant be empty")
	}

	if req.Password == "" {
		badRequestDetails = append(badRequestDetails, "Password cant be empty")
	}

	if len(badRequestDetails) != 0 {
		return badRequestDetails, errors.ErrBadRequest
	}

	return nil, nil
}
