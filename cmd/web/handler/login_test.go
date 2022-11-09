package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kompiang_mini-project_social-media/cmd/web/path"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestLogin() {
	tests := []struct {
		name             string
		request          dto.UserLoginRequest
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcLoginService
		contentType      string
	}{
		{
			name: "Valid",
			request: dto.UserLoginRequest{
				Username: "test",
				Password: "12345678",
			},
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.UserSuccessLoginResponse{
					AccessToken:  "123",
					RefreshToken: "123",
				},
			},
			service: func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
				return &dto.UserSuccessLoginResponse{
					AccessToken:  "123",
					RefreshToken: "123",
				}, nil
			},
			contentType: echo.MIMEApplicationJSON,
		},
		{
			name: "Invalid Content Type",
			request: dto.UserLoginRequest{
				Username: "test",
				Password: "12345678",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  "Content type must be application/json",
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
				return &dto.UserSuccessLoginResponse{
					AccessToken:  "123",
					RefreshToken: "123",
				}, nil
			},
			contentType: echo.MIMEApplicationProtobuf,
		},
		{
			name: "Invalid Empty Field",
			request: dto.UserLoginRequest{
				Username: "",
				Password: "12345678",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"Username cant be empty"},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
				return &dto.UserSuccessLoginResponse{
					AccessToken:  "123",
					RefreshToken: "123",
				}, nil
			},
			contentType: echo.MIMEApplicationJSON,
		},
		{
			name: "Invalid Empty Field",
			request: dto.UserLoginRequest{
				Username: "",
				Password: "",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail: []string{
						"Username cant be empty",
						"Password cant be empty",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
				return &dto.UserSuccessLoginResponse{
					AccessToken:  "123",
					RefreshToken: "123",
				}, nil
			},
			contentType: echo.MIMEApplicationJSON,
		},
		{
			name: "Service Return Error",
			request: dto.UserLoginRequest{
				Username: "test",
				Password: "12345678",
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
				return nil, errors.ErrInternalServer
			},
			contentType: echo.MIMEApplicationJSON,
		},
	}

	for _, test := range tests {
		request, err := json.Marshal(test.request)
		if err != nil {
			log.Fatal(err.Error())
		}

		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			log.Fatal(err.Error())
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, path.Register, strings.NewReader(string(request)))
		req.Header.Add(echo.HeaderContentType, test.contentType)

		ctx := s.E.NewContext(req, res)

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, Login(mockService{funcLoginService: test.service})(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
