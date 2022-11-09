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

func (s handlerSuite) TestRefreshToken() {
	tests := []struct {
		name             string
		request          dto.RefreshTokenRequest
		expectedResponse dto.BaseResponse
		expectedCode     int
		contentType      string
		serviceMock      funcRefreshTokenService
	}{
		{
			name: "Valid",
			request: dto.RefreshTokenRequest{
				RefreshToken: "refresh",
			},
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.RefreshTokenResponse{
					AccessToken: "access",
				},
			},
			expectedCode: http.StatusOK,
			contentType:  echo.MIMEApplicationJSON,
			serviceMock: func(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
				return &dto.RefreshTokenResponse{
					AccessToken: "access",
				}, nil
			},
		},
		{
			name: "Content type not valid",
			request: dto.RefreshTokenRequest{
				RefreshToken: "refresh",
			},
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  "Content type must be application/json",
				},
				Data: nil,
			},
			expectedCode: http.StatusBadRequest,
			contentType:  echo.MIMEApplicationProtobuf,
			serviceMock: func(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
				return &dto.RefreshTokenResponse{
					AccessToken: "access",
				}, nil
			},
		},
		{
			name: "Field request cant be empty",
			request: dto.RefreshTokenRequest{
				RefreshToken: "",
			},
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  "Refresh Token field couldn't be empty",
				},
				Data: nil,
			},
			expectedCode: http.StatusBadRequest,
			contentType:  echo.MIMEApplicationJSON,
			serviceMock: func(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
				return &dto.RefreshTokenResponse{
					AccessToken: "access",
				}, nil
			},
		},
		{
			name: "Service Return Error",
			request: dto.RefreshTokenRequest{
				RefreshToken: "refresh",
			},
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			expectedCode: http.StatusInternalServerError,
			contentType:  echo.MIMEApplicationJSON,
			serviceMock: func(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
				return nil, errors.ErrInternalServer
			},
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
			if assert.NoError(t, RefreshToken(mockService{funcRefreshTokenService: test.serviceMock})(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
