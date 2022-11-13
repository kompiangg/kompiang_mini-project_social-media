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
	"github.com/kompiang_mini-project_social-media/pkg/utils/stringutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestCreateComment() {
	tests := []struct {
		name    string
		request struct {
			content     string
			postID      string
			recommentID string
			image       *string
			video       *string
			userContext *dto.UserContext
		}
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcCreateComment
	}{
		{
			name: "Valid",
			request: struct {
				content     string
				postID      string
				recommentID string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				postID:      "valid",
				recommentID: "",
				image:       stringutils.StringToPointer(s.ImageAbsolutePath),
				video:       stringutils.StringToPointer(s.VideoAbsolutePath),
				userContext: s.UserContext,
			},
			expectedCode: http.StatusCreated,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.CommentResponse{
					ID:            "valid",
					CommentedBy:   "valid",
					Content:       "valid",
					PostID:        stringutils.StringToPointer("valid"),
					RecommentID:   nil,
					PictureURL:    stringutils.StringToPointer("https://cloudinary"),
					VideoURL:      stringutils.StringToPointer("https://cloudinary"),
					ChildComments: []string{},
					CreatedAt:     time.Time{},
				},
			},
			service: func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName, videoFileName *string) (*dto.CommentResponse, error) {
				return &dto.CommentResponse{
					ID:            "valid",
					CommentedBy:   "valid",
					Content:       "valid",
					PostID:        stringutils.StringToPointer("valid"),
					RecommentID:   nil,
					PictureURL:    stringutils.StringToPointer("https://cloudinary"),
					VideoURL:      stringutils.StringToPointer("https://cloudinary"),
					ChildComments: []string{},
					CreatedAt:     time.Time{},
				}, nil
			},
		},
		{
			name: "Post ID and Recomment ID empty",
			request: struct {
				content     string
				postID      string
				recommentID string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				postID:      "",
				recommentID: "",
				image:       stringutils.StringToPointer(s.ImageAbsolutePath),
				video:       stringutils.StringToPointer(s.VideoAbsolutePath),
				userContext: s.UserContext,
			},
			expectedCode: http.StatusBadRequest,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrBadRequest.Error(),
					Detail:  []string{"One of post_id or recomment_id shouldnt be empty"},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName, videoFileName *string) (*dto.CommentResponse, error) {
				return nil, nil
			},
		},
		{
			name: "Content Empty",
			request: struct {
				content     string
				postID      string
				recommentID string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "",
				postID:      "Valid",
				recommentID: "",
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
			service: func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName, videoFileName *string) (*dto.CommentResponse, error) {
				return nil, nil
			},
		},
		{
			name: "Cant extract user context",
			request: struct {
				content     string
				postID      string
				recommentID string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				postID:      "Valid",
				recommentID: "",
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
			service: func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName, videoFileName *string) (*dto.CommentResponse, error) {
				return nil, nil
			},
		},
		{
			name: "Service return error",
			request: struct {
				content     string
				postID      string
				recommentID string
				image       *string
				video       *string
				userContext *dto.UserContext
			}{
				content:     "Valid",
				postID:      "Valid",
				recommentID: "",
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
			service: func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName, videoFileName *string) (*dto.CommentResponse, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		requestBody := new(bytes.Buffer)
		writer := multipart.NewWriter(requestBody)
		writer.WriteField("post_id", test.request.postID)
		writer.WriteField("recomment_id", test.request.recommentID)
		writer.WriteField("content", test.request.content)
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

		req := httptest.NewRequest(http.MethodPost, "/api/v1/comment", requestBody)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.request.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, CreateComment(mockService{funcCreateComment: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}

func (s handlerSuite) TestGetChildCommentByCommentID() {
	tests := []struct {
		name             string
		parentCommentID  string
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetChildComment
	}{
		{
			name:            "Valid",
			parentCommentID: "Valid",
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: &dto.CommentsResponse{
					{
						ID:            "Valid",
						CommentedBy:   "Valid",
						Content:       "Valid",
						PostID:        stringutils.StringToPointer("Valid"),
						RecommentID:   stringutils.StringToPointer("Valid"),
						PictureURL:    nil,
						VideoURL:      nil,
						ChildComments: nil,
						CreatedAt:     time.Time{},
					},
				},
			},
			expectedCode: http.StatusOK,
			service: func(ctx context.Context, parentID string) (*dto.CommentsResponse, error) {
				return &dto.CommentsResponse{
					{
						ID:            "Valid",
						CommentedBy:   "Valid",
						Content:       "Valid",
						PostID:        stringutils.StringToPointer("Valid"),
						RecommentID:   stringutils.StringToPointer("Valid"),
						PictureURL:    nil,
						VideoURL:      nil,
						ChildComments: nil,
						CreatedAt:     time.Time{},
					},
				}, nil
			},
		},
		{
			name:            "Service return invalid",
			parentCommentID: "Valid",
			expectedCode:    http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, parentID string) (*dto.CommentsResponse, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodGet, "/api/v1/comment/:parent_comment_id", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.SetParamNames("parent_comment_id")
		c.SetParamValues(test.parentCommentID)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetChildCommentByCommentID(mockService{funcGetChildComment: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}
