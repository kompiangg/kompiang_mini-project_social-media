package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestDeactivateUser() {
	tests := []struct {
		name             string
		userContext      *dto.UserContext
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcDeactivateAccount
	}{
		{
			name:         "Valid",
			userContext:  s.UserContext,
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data:  "User Deactivate!",
			},
			service: func(ctx context.Context, username string) error {
				return nil
			},
		},
		{
			name:         "Failed extract user context",
			userContext:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, username string) error {
				return nil
			},
		},
		{
			name:         "Servie return error",
			userContext:  s.UserContext,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, username string) error {
				return errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/deactivate", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, DeactivateUser(mockService{funcDeactivateAccount: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestEditProfile() {
	tests := []struct {
		name             string
		request          dto.EditProfileRequest
		userContext      *dto.UserContext
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcEditUserService
	}{
		{
			name: "Valid",
			request: dto.EditProfileRequest{
				Username:          "Valid",
				Email:             "Valid",
				DisplayName:       "Valid",
				ProfilePictureURL: nil,
				Bio:               "Valid",
				BirthDate:         "1970-01-01",
			},
			userContext:  s.UserContext,
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.EditProfileResponse{
					Username:          "Valid",
					Email:             "Valid",
					DisplayName:       "Valid",
					ProfilePictureURL: "",
					Bio:               "Valid",
					BirthDate:         &time.Time{},
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				},
			},
			service: func(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) {
				return &dto.EditProfileResponse{
					Username:          "Valid",
					Email:             "Valid",
					DisplayName:       "Valid",
					ProfilePictureURL: "",
					Bio:               "Valid",
					BirthDate:         &time.Time{},
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil
			},
		},
		{
			name: "Failed to extract user context",
			request: dto.EditProfileRequest{
				Username:          "Valid",
				Email:             "Valid",
				DisplayName:       "Valid",
				ProfilePictureURL: nil,
				Bio:               "Valid",
				BirthDate:         "1970-01-01",
			},
			userContext:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) {
				return nil, nil
			},
		},
		{
			name: "Service return error",
			request: dto.EditProfileRequest{
				Username:          "Valid",
				Email:             "Valid",
				DisplayName:       "Valid",
				ProfilePictureURL: nil,
				Bio:               "Valid",
				BirthDate:         "1970-01-01",
			},
			userContext:  s.UserContext,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		requestBody := new(bytes.Buffer)
		writer := multipart.NewWriter(requestBody)
		writer.WriteField("email", test.request.Email)
		writer.WriteField("display_name", test.request.DisplayName)
		writer.WriteField("bio", test.request.Bio)
		writer.WriteField("birth_date", test.request.BirthDate)

		if test.request.ProfilePictureURL != nil {
			writer.CreateFormFile("image", *test.request.ProfilePictureURL)
		}

		err := writer.Close()
		if err != nil {
			panic(err)
		}

		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/comment", requestBody)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, EditProfile(mockService{funcEditUserService: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestGetProfile() {
	tests := []struct {
		name                string
		params              string
		userContext         *dto.UserContext
		serviceMyProfile    funcGetMyProfile
		serviceOtherProfile funcGetOtherUserProfile
		expectedResponse    dto.BaseResponse
		expectedCode        int
	}{
		{
			name:        "Valid My Profile",
			params:      s.UserContext.Username,
			userContext: s.UserContext,
			serviceMyProfile: func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
				return &dto.UserMyProfileResponse{
					Username:          "Valid",
					Email:             "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					IsVerifiedEmail:   false,
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil
			},
			serviceOtherProfile: func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
				return &dto.UserOtherProfileResponse{
					Username:          "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil
			},
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.UserMyProfileResponse{
					Username:          "Valid",
					Email:             "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					IsVerifiedEmail:   false,
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			name:        "Valid Other Profile",
			params:      "TestUser",
			userContext: s.UserContext,
			serviceMyProfile: func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
				return &dto.UserMyProfileResponse{
					Username:          "Valid",
					Email:             "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					IsVerifiedEmail:   false,
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil
			},
			serviceOtherProfile: func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
				return &dto.UserOtherProfileResponse{
					Username:          "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil
			},
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.UserOtherProfileResponse{
					Username:          "Valid",
					CountFollowers:    0,
					CountFollowings:   0,
					DisplayName:       "Valid",
					ProfilePictureURL: "Valid",
					Bio:               "Valid",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			name:        "Failed to extract user context",
			params:      s.UserContext.Username,
			userContext: nil,
			serviceMyProfile: func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
				return nil, nil
			},
			serviceOtherProfile: func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
				return nil, nil
			},
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:        "My profile service return error",
			params:      s.UserContext.Username,
			userContext: s.UserContext,
			serviceMyProfile: func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
				return nil, errors.ErrInternalServer
			},
			serviceOtherProfile: func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
				return nil, nil
			},
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:        "Other profile service return error",
			params:      "TestUser",
			userContext: s.UserContext,
			serviceMyProfile: func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
				return nil, nil
			},
			serviceOtherProfile: func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
				return nil, errors.ErrInternalServer
			},
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
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/:username", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("username")
		c.SetParamValues(test.params)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetProfile(mockService{funcGetMyProfile: test.serviceMyProfile, funcGetOtherUserProfile: test.serviceOtherProfile})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}
