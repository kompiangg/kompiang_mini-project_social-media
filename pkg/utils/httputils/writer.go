package httputils

import (
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"

	"github.com/labstack/echo/v4"
)

type SuccessResponseParams struct {
	Code int
	Data interface{}
}

type ErrorResponseParams struct {
	Err    error
	Detail interface{}
}

func WriteResponse(c echo.Context, data SuccessResponseParams) error {
	if data.Code == 0 {
		data.Code = 200
	}
	return c.JSON(data.Code, dto.BaseResponse{
		Error: nil,
		Data:  data.Data,
	})
}

func WriteErrorResponse(c echo.Context, params ErrorResponseParams) error {
	e := errors.GetErr(params.Err)
	return c.JSON(e.HTTPErrorCode, dto.BaseResponse{
		Error: &dto.ErrorBaseResponse{
			Message: e.Message,
			Detail:  params.Detail,
		},
		Data: nil,
	})
}
