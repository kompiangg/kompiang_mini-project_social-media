package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestGetFollowers() {
	tests := []struct {
		name             string
		params           string
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetFollowers
	}{
		{
			name:         "Valid",
			params:       s.UserContext.Username,
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: &dto.Followers{
					{
						Username: "pang",
					},
				},
			},
			service: func(ctx context.Context, username string) (*dto.Followers, error) {
				return &dto.Followers{
					{
						Username: "pang",
					},
				}, nil
			},
		},
		{
			name:         "Service Return error",
			params:       s.UserContext.Username,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, username string) (*dto.Followers, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodGet, "/api/v1/followers/:username", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("username")
		c.SetParamValues(test.params)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetFollowers(mockService{funcGetFollowers: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestGetFollowing() {
	tests := []struct {
		name             string
		params           string
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetFollowings
	}{
		{
			name:         "Valid",
			params:       s.UserContext.Username,
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: &dto.Followings{
					{
						Username: "pang",
					},
				},
			},
			service: func(ctx context.Context, username string) (*dto.Followings, error) {
				return &dto.Followings{
					{
						Username: "pang",
					},
				}, nil
			},
		},
		{
			name:         "Service Return error",
			params:       s.UserContext.Username,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, username string) (*dto.Followings, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodGet, "/api/v1/followers/:username", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("username")
		c.SetParamValues(test.params)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetFollowing(mockService{funcGetFollowings: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestFollowOtherUser() {
	tests := []struct {
		name             string
		requestBody      dto.FollowRequest
		userContext      *dto.UserContext
		service          funcFollowOtherUser
		contentType      string
		expectedResponse dto.BaseResponse
		expectedCode     int
	}{
		{
			name: "Valid",
			requestBody: dto.FollowRequest{
				FollowRequestToUser: "Valid",
			},
			userContext: s.UserContext,
			service: func(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
				return nil
			},
			contentType: echo.MIMEApplicationJSON,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data:  "Followed",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Failed binding data",
			requestBody: dto.FollowRequest{
				FollowRequestToUser: "Valid",
			},
			userContext: s.UserContext,
			service: func(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
				return nil
			},
			contentType: echo.MIMEApplicationProtobuf,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"Content type must be application/json"},
				},
				Data: nil,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Service return error",
			requestBody: dto.FollowRequest{
				FollowRequestToUser: "Valid",
			},
			userContext: s.UserContext,
			service: func(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
				return errors.ErrInternalServer
			},
			contentType: echo.MIMEApplicationJSON,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		requestBody, err := json.Marshal(test.requestBody)
		if err != nil {
			panic(err)
		}

		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/follows", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, test.contentType)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, FollowOtherUser(mockService{funcFollowOtherUser: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}
