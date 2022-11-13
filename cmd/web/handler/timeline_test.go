package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestGetTimeline() {
	tests := []struct {
		name             string
		userContext      *dto.UserContext
		expectedCode     int
		expectedResponse dto.BaseResponse
		service          funcGetTimeline
	}{
		{
			name:         "Valid",
			userContext:  s.UserContext,
			expectedCode: http.StatusOK,
			expectedResponse: dto.BaseResponse{
				Error: nil,
				Data: dto.Timeline{
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
			service: func(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
				return &dto.Timeline{
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
			name:         "Cant extract user context",
			userContext:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
				return nil, nil
			},
		},
		{
			name:         "Service return error",
			userContext:  s.UserContext,
			expectedCode: http.StatusInternalServerError,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
				return nil, errors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/timeline", nil)
		res := httptest.NewRecorder()
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, GetTimeline(mockService{funcGetTimeline: test.service})(c)) {
				assert.Equal(t, test.expectedCode, res.Code)
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
			}
		})
	}
}
