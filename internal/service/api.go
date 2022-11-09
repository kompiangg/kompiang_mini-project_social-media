package service

import (
	"context"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

type Service interface {
	// Auth
	RegisterService(ctx context.Context, req *dto.UserRegisterRequest) error                                           // Finish
	LoginService(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error)                // Finish
	RefreshTokenService(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) // Finish

	// User
	EditUserService(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) // Finish
	DeactivateUser(ctx context.Context, username string) error                                                                                                    // Finish
	GetOtherUserProfile(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error)                                                              // Finish
	GetMyProfile(ctx context.Context, username string) (*dto.UserMyProfileResponse, error)                                                                        // Finish

	// Posts
	CreatePost(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName *string, videoFileName *string) (*dto.Post, error) // Finish
	DeletePost(ctx context.Context, req *dto.DeletePostRequest) error                                                                             // Finish
	GetPostByID(ctx context.Context, id string) (*dto.Post, error)                                                                                // Finish
	GetAllPostByUsername(ctx context.Context, username string) (*dto.Posts, error)                                                                // Finish

	// Comment
	CreateComment(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName *string, videoFileName *string) (*dto.CommentResponse, error) // Finish
	GetChildComment(ctx context.Context, parentID string) (*dto.CommentsResponse, error)                                                                           // Finish

	// Follow (User Relation)
	FollowOtherUser(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error // Finish
	GetFollowers(ctx context.Context, username string) (*dto.Followers, error)                   // Finish
	GetFollowings(ctx context.Context, username string) (*dto.Followings, error)                 // Finish

	// Timeline
	GetTimeline(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error)

	// Notifications
	CreateGeneralNotifications(ctx context.Context, req *dto.AdminNotificationRequest) error
	GetNotificationsByUsername(ctx context.Context, username string) (*dto.NotificationsResponse, error)
}
