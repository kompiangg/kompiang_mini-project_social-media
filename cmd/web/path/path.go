package path

import "github.com/labstack/echo/v4"

const (
	apiV1GroupPath = "/api/v1"

	pingRootGroupPath = "/ping"
	// Auth
	authRootGroupPath = "/auth"
	// User Profile
	userRootGroupPath = "/user"
	// Post
	postRootGroupPath = "/post"
	// Comment
	commentRootGroupPath = "/comment"
	// Timeline
	timelineRootGroupPath = "/timeline"
	// Notifications
	notificationRootGroupPath = "/notification"
	// User Relation
	userRelationRootGroupPath = ""
	// WebSocket
	WebSocketGroupRootGroupPath = "/ws"
)

type PathGroups struct {
	ApiGroupV1          *echo.Group
	PingGroupV1         *echo.Group
	AuthGroupV1         *echo.Group
	UserProfileGroupV1  *echo.Group
	PostGroupV1         *echo.Group
	CommentGroupV1      *echo.Group
	TimelineGroupV1     *echo.Group
	NotificationGroupV1 *echo.Group
	UserRelationGroupV1 *echo.Group
	WebSocketGroup      *echo.Group
}

func InitPathGroups(e *echo.Echo) *PathGroups {
	pathGroups := PathGroups{}

	pathGroups.ApiGroupV1 = e.Group(apiV1GroupPath)
	pathGroups.PingGroupV1 = pathGroups.ApiGroupV1.Group(pingRootGroupPath)
	pathGroups.AuthGroupV1 = pathGroups.ApiGroupV1.Group(authRootGroupPath)
	pathGroups.UserProfileGroupV1 = pathGroups.ApiGroupV1.Group(userRootGroupPath)
	pathGroups.PostGroupV1 = pathGroups.ApiGroupV1.Group(postRootGroupPath)
	pathGroups.CommentGroupV1 = pathGroups.ApiGroupV1.Group(commentRootGroupPath)
	pathGroups.TimelineGroupV1 = pathGroups.ApiGroupV1.Group(timelineRootGroupPath)
	pathGroups.NotificationGroupV1 = pathGroups.ApiGroupV1.Group(notificationRootGroupPath)
	pathGroups.UserRelationGroupV1 = pathGroups.ApiGroupV1.Group(userRelationRootGroupPath)
	pathGroups.WebSocketGroup = pathGroups.ApiGroupV1.Group(WebSocketGroupRootGroupPath)

	return &pathGroups
}
