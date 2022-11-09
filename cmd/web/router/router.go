package router

import (
	"github.com/kompiang_mini-project_social-media/cmd/web/handler"
	"github.com/kompiang_mini-project_social-media/cmd/web/path"
	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/middleware"
	"github.com/kompiang_mini-project_social-media/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
)

type RouteParams struct {
	E       *echo.Echo
	Service service.Service
	Config  *config.Config
	Pool    *websocketutils.Pool
}

func InitRoute(params RouteParams) {
	params.E.GET(path.Ping, handler.PingHandler())

	// Auth
	params.E.POST(path.Register, handler.Register(params.Service))
	params.E.POST(path.Login, handler.Login(params.Service))
	params.E.POST(path.RefreshToken, handler.RefreshToken(params.Service))
	// User Profile
	params.E.PUT(path.DeactivateUser, handler.DeactivateUser(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.GET(path.GetProfileByUsername, handler.GetProfile(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.PUT(path.EditProfile, handler.EditProfile(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	// Post
	params.E.POST(path.CreatePost, handler.CreatePost(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.DELETE(path.DeletePost, handler.DeletePost(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.GET(path.GetPostByID, handler.GetPostById(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.GET(path.GetPostByUsername, handler.GetPostByUsername(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	// Comment
	params.E.POST(path.CreateComment, handler.CreateComment(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.GET(path.GetChildCommentByParentID, handler.GetChildCommentByCommentID(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	// Timeline
	params.E.GET(path.Timeline, handler.GetTimeline(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	// Notifications
	params.E.GET(path.GetNotifications, handler.GetNotifications(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.POST(path.CreateGeneralNotifications, handler.AdminCreateGeneralNotifications(params.Service), middleware.AdminSecretTokenAuth(&params.Config.Admin), middleware.AdminBasicAuth(&params.Config.Admin))
	// User Relation
	params.E.GET(path.GetFollowers, handler.GetFollowers(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.GET(path.GetFollowings, handler.GetFollowing(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	params.E.POST(path.FollowOtherUser, handler.FollowOtherUser(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))

	params.E.GET(path.WebSocket, handler.WebSocket(params.Pool), middleware.MustAuthorized(&params.Config.JWTConfig))
}
