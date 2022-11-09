package handler

import (
	"log"
	"net/http"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func Register(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var registerReq dto.UserRegisterRequest
		if err := c.Bind(&registerReq); err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: "Content type must be application/json",
			})
		}

		detail, err := validateRegister(registerReq)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    err,
				Detail: detail,
			})
		}

		err = service.RegisterService(c.Request().Context(), &registerReq)
		if err == errors.ErrAccountDuplicated {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrAccountDuplicated,
			})
		}

		if err != nil {
			log.Println("[HANDLER ERROR] While calling the register service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: http.StatusCreated,
			Data: "Account Created",
		})
	}
}

func validateRegister(req dto.UserRegisterRequest) ([]string, error) {
	var badRequestDetails []string

	if req.Username == "" {
		badRequestDetails = append(badRequestDetails, "Username cant be empty")
	} else if len(req.Username) >= 50 {
		badRequestDetails = append(badRequestDetails, "Username cant not longer than 50")
	}

	if req.Email == "" {
		badRequestDetails = append(badRequestDetails, "Email cant be empty")
	}

	if req.Password == "" {
		badRequestDetails = append(badRequestDetails, "Password cant be empty")
	}

	if req.DisplayName == "" {
		badRequestDetails = append(badRequestDetails, "Display name cant be empty")
	}

	if len(badRequestDetails) != 0 {
		return badRequestDetails, errors.ErrBadRequest
	}

	return nil, nil
}
