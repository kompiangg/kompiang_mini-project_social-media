package path

const (
	Ping = "/api/v1/ping"
	// Auth
	Register     = "/api/v1/auth/register"
	Login        = "/api/v1/auth/login"
	RefreshToken = "/api/v1/auth/refresh"
	// User Profile
	DeactivateUser       = "/api/v1/user/deactivate"
	GetProfileByUsername = "/api/v1/user/:username"
	EditProfile          = "/api/v1/user"
	GetPostByUsername    = "/api/v1/user/:username/post"
	// Post
	CreatePost  = "/api/v1/post"
	DeletePost  = "/api/v1/post"
	GetPostByID = "/api/v1/post/:id"
	// Comment
	CreateComment             = "/api/v1/comment"
	GetChildCommentByParentID = "/api/v1/comment/:parent_comment_id"
	// Timeline
	Timeline = "/api/v1/timeline"
	// Notifications
	CreateGeneralNotifications = "/api/v1/notifications"
	GetNotifications           = "/api/v1/notifications"
	// User Relation
	GetFollowers    = "/api/v1/followers/:username"
	GetFollowings   = "/api/v1/followings/:username"
	FollowOtherUser = "/api/v1/follow"
	// WebSocket
	WebSocket = "/ws"
)
