package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kompiang_mini-project_social-media/cmd/web/path"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func (s handlerSuite) TestPing() {
	expectedResponse := dto.BaseResponse{
		Error: nil,
		Data:  "pong",
	}

	responseStr, err := json.Marshal(expectedResponse)
	if err != nil {
		log.Fatal(err.Error())
	}

	expectedCode := 200

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path.Ping, nil)

	ctx := s.E.NewContext(req, res)

	s.Suite.T().Run("Test Ping", func(t *testing.T) {
		if assert.NoError(t, PingHandler()(ctx)) {
			assert.Equal(t, string(responseStr), res.Body.String()[:len(res.Body.String())-1])
			assert.Equal(t, expectedCode, res.Code)
		}
	})
}
