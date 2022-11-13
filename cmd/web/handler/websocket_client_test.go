package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestWebSocket() {
	tests := []struct {
		name             string
		userContext      *dto.UserContext
		expectedResponse interface{}
		upgrade          string
		connection       string
	}{
		{
			name:        "Cant extract user context",
			userContext: nil,
			expectedResponse: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: errors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			upgrade:    "websocket",
			connection: "Upgrade",
		},
		{
			name:             "Client not using websocket protocol",
			userContext:      s.UserContext,
			expectedResponse: "Bad Request",
		},
		{
			name:             "Valid",
			userContext:      s.UserContext,
			expectedResponse: "Bad Request",
			upgrade:          "websocket",
			connection:       "Upgrade",
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponse)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/ws", nil)
		res := httptest.NewRecorder()
		res.Header().Set(echo.HeaderConnection, test.connection)
		res.Header().Set(echo.HeaderUpgrade, test.upgrade)
		c := s.E.NewContext(req, res)
		c.Set(dto.AccountCTXKey, test.userContext)

		s.Pool = websocketutils.NewPool()

		go func() {
			s.Pool.Start()
		}()

		s.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, WebSocket(s.Pool)(c)) {
				time.Sleep(time.Second * 1)
				if res.Header().Get(echo.HeaderContentType) == "text/plain; charset=utf-8" {
					responseData, _ := io.ReadAll(res.Body)
					log.Println(string(responseData))
					assert.Equal(t, string(expectedResponse), fmt.Sprintf("\"%s\"", string(responseData)[:len(responseData)-1]))
				} else {
					assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				}
			}
		})
	}
}
