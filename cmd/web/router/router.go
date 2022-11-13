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
	pathGroups := path.InitPathGroups(params.E)

	pathGroups.PingGroupV1.GET("", handler.PingHandler())
	// Auth => /api/v1/auth
	pathGroups.AuthGroupV1.POST("/register", handler.Register(params.Service))
	pathGroups.AuthGroupV1.POST("/login", handler.Login(params.Service))
	pathGroups.AuthGroupV1.POST("/refresh", handler.RefreshToken(params.Service))
	// User Profile => /api/v1/user
	pathGroups.UserProfileGroupV1.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.UserProfileGroupV1.PUT("/deactivate", handler.DeactivateUser(params.Service))
	pathGroups.UserProfileGroupV1.PUT("", handler.EditProfile(params.Service))
	pathGroups.UserProfileGroupV1.GET("/:username", handler.GetProfile(params.Service))
	// Post => /api/v1/post
	pathGroups.PostGroupV1.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.PostGroupV1.POST("", handler.CreatePost(params.Service))
	pathGroups.PostGroupV1.DELETE("", handler.DeletePost(params.Service))
	pathGroups.PostGroupV1.GET("/:id", handler.GetPostById(params.Service))
	pathGroups.PostGroupV1.GET("/:username/post", handler.GetPostByUsername(params.Service))
	// Comment => /api/v1/comment
	pathGroups.CommentGroupV1.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.CommentGroupV1.POST("", handler.CreateComment(params.Service))
	pathGroups.CommentGroupV1.GET("/:parent_comment_id", handler.GetChildCommentByCommentID(params.Service))
	// Timeline => /api/v1/timeline
	pathGroups.TimelineGroupV1.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.TimelineGroupV1.GET("", handler.GetTimeline(params.Service))
	// Notifications => /api/v1/notifications
	pathGroups.NotificationGroupV1.GET("/notifications", handler.GetNotifications(params.Service), middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.NotificationGroupV1.POST("/notifications", handler.AdminCreateGeneralNotifications(params.Service), middleware.AdminSecretTokenAuth(&params.Config.Admin), middleware.AdminBasicAuth(&params.Config.Admin))
	// User Relation => /api/v1
	pathGroups.UserRelationGroupV1.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.UserRelationGroupV1.GET("/followers/:username", handler.GetFollowers(params.Service))
	pathGroups.UserRelationGroupV1.GET("/followings/:username", handler.GetFollowing(params.Service))
	pathGroups.UserRelationGroupV1.POST("/follow", handler.FollowOtherUser(params.Service))
	// Websocket => /api/v1/ws
	pathGroups.WebSocketGroup.Use(middleware.MustAuthorized(&params.Config.JWTConfig))
	pathGroups.WebSocketGroup.GET("", handler.WebSocket(params.Pool))
}
