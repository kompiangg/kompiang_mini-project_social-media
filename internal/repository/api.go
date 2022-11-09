package repository

import (
	"context"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

type Repository interface {
	// Auth
	AccountLogin(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserLogin, error)     // Finish
	InsertRefreshToken(ctx context.Context, refreshToken *dto.UserRefreshToken) error        // Finish
	RegisterUser(ctx context.Context, req *dto.UserRegisterRequest) error                    // Finish
	GetRefreshToken(ctx context.Context, refreshToken string) (*dto.UserRefreshToken, error) // Finish
	GetUserByUsername(ctx context.Context, username string) (*dto.UserContext, error)        // Finish

	// User
	EditUser(ctx context.Context, req *dto.EditProfileRequest, userCtx *dto.UserContext) (*dto.EditProfileResponse, error)
	DeactivateUser(ctx context.Context, username string) error
	GetOtherUserProfile(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error)
	GetMyProfile(ctx context.Context, username string) (*dto.UserMyProfileResponse, error)

	// Post
	CreatePost(ctx context.Context, req *dto.CreatePostRequest, username string, id string) (*dto.Post, error)
	DeletePost(ctx context.Context, req *dto.DeletePostRequest) error
	GetPostByID(ctx context.Context, id string) (*dto.Post, error)
	GetAllPostByUsername(ctx context.Context, username string) (*dto.Posts, error)

	// Comment
	CreateComment(ctx context.Context, req *dto.CreateCommentRequest, username string, id string) (*dto.CommentResponse, error)
	GetChildComment(ctx context.Context, parentID string) (*dto.CommentsResponse, error)

	// Follow (User Relation)
	FollowOtherUser(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error
	GetFollowers(ctx context.Context, username string) (*dto.Followers, error)
	GetFollowings(ctx context.Context, username string) (*dto.Followings, error)
	// Timeline
	GetTimeline(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error)
	CreateGeneralNotifications(ctx context.Context, req *dto.AdminNotificationRequest) error
	GetNotificationsByUsername(ctx context.Context, username string) (*dto.NotificationsResponse, error)
}
