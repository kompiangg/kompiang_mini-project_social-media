package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/stringutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestCreatePost() {
	tests := []struct {
		name    string
		request struct {
			content     string
			repostID    *string
			image       *string
			video       *string
			userContext *dto.UserContext
		}
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcCreatePost
	}{
		{
			name: "Valid",
			request: struct {
				content     string
				repostID    *string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				repostID:    nil,
				image:       nil,
				video:       nil,
				userContext: s.UserContext,
			},
			expectedCode: http.StatusCreated,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: &dto.Post{
					ID:          "Valid",
					PublishedBy: "Valid",
					Content:     "Valid",
					RepostID:    stringutils.StringToPointer("Valid"),
					PictureURL:  nil,
					VideoURL:    nil,
					Comments:    nil,
					CreatedAt:   time.Time{},
				},
			},
			service: func(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName, videoFileName *string) (*dto.Post, error) {
				return &dto.Post{
					ID:          "Valid",
					PublishedBy: "Valid",
					Content:     "Valid",
					RepostID:    stringutils.StringToPointer("Valid"),
					PictureURL:  nil,
					VideoURL:    nil,
					Comments:    nil,
					CreatedAt:   time.Time{},
				}, nil
			},
		},
		{
			name: "Cant extract yser request",
			request: struct {
				content     string
				repostID    *string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				repostID:    nil,
				image:       nil,
				video:       nil,
				userContext: nil,
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName, videoFileName *string) (*dto.Post, error) {
				return nil, nil
			},
		},
		{
			name: "Content invalid",
			request: struct {
				content     string
				repostID    *string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "",
				repostID:    nil,
				image:       nil,
				video:       nil,
				userContext: s.UserContext,
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"Content field can't be empty"},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName, videoFileName *string) (*dto.Post, error) {
				return nil, nil
			},
		},
		{
			name: "Service return error",
			request: struct {
				content     string
				repostID    *string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				repostID:    nil,
				image:       nil,
				video:       nil,
				userContext: s.UserContext,
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName, videoFileName *string) (*dto.Post, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		requestBody := new(bytes.Buffer)
		writer := multipart.NewWriter(requestBody)
		writer.WriteField("content", test.request.content)
		if test.request.repostID != nil {
			writer.WriteField("repost_id", *test.request.repostID)
		}
		if test.request.image != nil {
			writer.CreateFormFile("image", *test.request.image)
		}
		if test.request.video != nil {
			writer.CreateFormFile("video", *test.request.image)
		}
		err := writer.Close()
		if err != nil {
			panic(err)
		}

		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/post", requestBody)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.request.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, CreatePost(mockService{funcCreatePost: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestDeletePost() {
	tests := []struct {
		name    string
		request struct {
			body        dto.DeletePostRequest
			userContext *dto.UserContext
			contentType string
		}
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcDeletePost
	}{
		{
			name: "Valid",
			request: struct {
				body        dto.DeletePostRequest
				userContext *dto.UserContext
				contentType string
			}{
				body: dto.DeletePostRequest{
					PostID: "Valid",
				},
				userContext: s.UserContext,
				contentType: echo.MIMEApplicationJSON,
			},
			service: func(ctx context.Context, req *dto.DeletePostRequest) error {
				return nil
			},
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data:  "Post Deleted!",
			},
		},
		{
			name: "Invalid request body",
			request: struct {
				body        dto.DeletePostRequest
				userContext *dto.UserContext
				contentType string
			}{
				body: dto.DeletePostRequest{
					PostID: "",
				},
				userContext: s.UserContext,
				contentType: echo.MIMEApplicationJSON,
			},
			service: func(ctx context.Context, req *dto.DeletePostRequest) error {
				return nil
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"post_id field can't be empty"},
				},
				Data: nil,
			},
		},
		{
			name: "Invalid content type",
			request: struct {
				body        dto.DeletePostRequest
				userContext *dto.UserContext
				contentType string
			}{
				body: dto.DeletePostRequest{
					PostID: "Valid",
				},
				userContext: s.UserContext,
				contentType: echo.MIMEMultipartForm,
			},
			service: func(ctx context.Context, req *dto.DeletePostRequest) error {
				return nil
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"Content type must be application/json"},
				},
				Data: nil,
			},
		},
		{
			name: "Service return error",
			request: struct {
				body        dto.DeletePostRequest
				userContext *dto.UserContext
				contentType string
			}{
				body: dto.DeletePostRequest{
					PostID: "Valid",
				},
				userContext: s.UserContext,
				contentType: echo.MIMEApplicationJSON,
			},
			service: func(ctx context.Context, req *dto.DeletePostRequest) error {
				return errors.ErrInternalServer
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		requestBody, err := json.Marshal(test.request.body)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/post", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, test.request.contentType)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.request.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, DeletePost(mockService{funcDeletePost: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestGetPostById() {
	tests := []struct {
		name             string
		params           string
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetPostByID
	}{
		{
			name:         "Valid",
			params:       "fdasfdaf",
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.Post{
					ID:          "Valid",
					PublishedBy: "Valid",
					Content:     "Valid",
					RepostID:    nil,
					PictureURL:  nil,
					VideoURL:    nil,
					Comments:    nil,
					CreatedAt:   time.Time{},
				},
			},
			service: func(ctx context.Context, id string) (*dto.Post, error) {
				return &dto.Post{
					ID:          "Valid",
					PublishedBy: "Valid",
					Content:     "Valid",
					RepostID:    nil,
					PictureURL:  nil,
					VideoURL:    nil,
					Comments:    nil,
					CreatedAt:   time.Time{},
				}, nil
			},
		},
		{
			name:         "Service return error",
			params:       "fdasfdaf",
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, id string) (*dto.Post, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/post/:id", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues(test.params)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetPostById(mockService{funcGetPostByID: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestGetPostByUsername() {
	tests := []struct {
		name             string
		params           string
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetAllPostByUsername
	}{
		{
			name:         "Valid",
			params:       "fdasfdaf",
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.Posts{
					{
						ID:          "Valid",
						PublishedBy: "Valid",
						Content:     "Valid",
						RepostID:    nil,
						PictureURL:  nil,
						VideoURL:    nil,
						Comments:    nil,
						CreatedAt:   time.Time{},
					},
				},
			},
			service: func(ctx context.Context, username string) (*dto.Posts, error) {
				return &dto.Posts{
					{
						ID:          "Valid",
						PublishedBy: "Valid",
						Content:     "Valid",
						RepostID:    nil,
						PictureURL:  nil,
						VideoURL:    nil,
						Comments:    nil,
						CreatedAt:   time.Time{},
					},
				}, nil
			},
		},
		{
			name:         "Service return error",
			params:       "fdasfdaf",
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, username string) (*dto.Posts, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/post/:id", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues(test.params)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetPostByUsername(mockService{funcGetAllPostByUsername: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}
