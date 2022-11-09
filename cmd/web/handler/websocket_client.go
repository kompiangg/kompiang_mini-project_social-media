package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebSocket(pool *websocketutils.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		wsConn, err := upgrade.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		client := &websocketutils.Client{
			Username: userCtx.Username,
			Conn:     wsConn,
			Pool:     pool,
		}

		pool.Register <- client
		client.Read()
		return nil
	}
}
