package handler

import (
	"context"
	"path"
	"path/filepath"
	"testing"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	E                 *echo.Echo
	Pool              *websocketutils.Pool
	UserContext       *dto.UserContext
	ImageAbsolutePath string
	VideoAbsolutePath string
}

func (s *handlerSuite) SetupSuite() {
	s.E = echo.New()
	s.UserContext = &dto.UserContext{
		Username:    "userTest",
		Email:       "userTest",
		DisplayName: "userTest",
	}
	var err error
	s.ImageAbsolutePath, err = filepath.Abs(path.Join("..", "..", "..", "pkg", "utils", "testutils", "test.png"))
	if err != nil {
		panic(err)
	}
	s.VideoAbsolutePath, err = filepath.Abs(path.Join("..", "..", "..", "pkg", "utils", "testutils", "test.mp4"))
	if err != nil {
		panic(err)
	}
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

type mockService struct {
	funcRegisterService
	funcLoginService
	funcRefreshTokenService
	funcEditUserService
	funcDeactivateAccount
	funcGetOtherUserProfile
	funcGetMyProfile
	funcCreatePost
	funcDeletePost
	funcGetPostByID
	funcGetAllPostByUsername
	funcCreateComment
	funcGetChildComment
	funcFollowOtherUser
	funcGetFollowers
	funcGetFollowings
	funcGetTimeline
	funcCreateGeneralNotifications
	funcGetNotificationsByUsername
}

type funcRegisterService func(ctx context.Context, req *dto.UserRegisterRequest) error
type funcLoginService func(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error)
type funcRefreshTokenService func(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
type funcEditUserService func(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error)
type funcDeactivateAccount func(ctx context.Context, username string) error
type funcGetOtherUserProfile func(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error)
type funcGetMyProfile func(ctx context.Context, username string) (*dto.UserMyProfileResponse, error)
type funcCreatePost func(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName *string, videoFileName *string) (*dto.Post, error)
type funcDeletePost func(ctx context.Context, req *dto.DeletePostRequest) error
type funcGetPostByID func(ctx context.Context, id string) (*dto.Post, error)
type funcGetAllPostByUsername func(ctx context.Context, username string) (*dto.Posts, error)
type funcCreateComment func(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName *string, videoFileName *string) (*dto.CommentResponse, error)
type funcGetChildComment func(ctx context.Context, parentID string) (*dto.CommentsResponse, error)
type funcFollowOtherUser func(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error
type funcGetFollowers func(ctx context.Context, username string) (*dto.Followers, error)
type funcGetFollowings func(ctx context.Context, username string) (*dto.Followings, error)
type funcGetTimeline func(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error)
type funcCreateGeneralNotifications func(ctx context.Context, req *dto.AdminNotificationRequest) error
type funcGetNotificationsByUsername func(ctx context.Context, username string) (*dto.NotificationsResponse, error)

func (s mockService) RegisterService(ctx context.Context, req *dto.UserRegisterRequest) error {
	return s.funcRegisterService(ctx, req)
}
func (s mockService) LoginService(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
	return s.funcLoginService(ctx, req)
}
func (s mockService) RefreshTokenService(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	return s.funcRefreshTokenService(ctx, refreshToken)
}
func (s mockService) EditUserService(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) {
	return s.funcEditUserService(ctx, userCtx, req, profilePictureFileName)
}
func (s mockService) DeactivateUser(ctx context.Context, username string) error {
	return s.funcDeactivateAccount(ctx, username)
}
func (s mockService) GetOtherUserProfile(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
	return s.funcGetOtherUserProfile(ctx, username)
}
func (s mockService) GetMyProfile(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
	return s.funcGetMyProfile(ctx, username)
}
func (s mockService) CreatePost(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName *string, videoFileName *string) (*dto.Post, error) {
	return s.funcCreatePost(ctx, req, username, imageFileName, videoFileName)
}
func (s mockService) DeletePost(ctx context.Context, req *dto.DeletePostRequest) error {
	return s.funcDeletePost(ctx, req)
}
func (s mockService) GetPostByID(ctx context.Context, id string) (*dto.Post, error) {
	return s.funcGetPostByID(ctx, id)
}
func (s mockService) GetAllPostByUsername(ctx context.Context, username string) (*dto.Posts, error) {
	return s.funcGetAllPostByUsername(ctx, username)
}
func (s mockService) CreateComment(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName *string, videoFileName *string) (*dto.CommentResponse, error) {
	return s.funcCreateComment(ctx, req, username, imageFileName, videoFileName)
}
func (s mockService) GetChildComment(ctx context.Context, parentID string) (*dto.CommentsResponse, error) {
	return s.funcGetChildComment(ctx, parentID)
}
func (s mockService) FollowOtherUser(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
	return s.funcFollowOtherUser(ctx, req, userCtx)
}
func (s mockService) GetFollowers(ctx context.Context, username string) (*dto.Followers, error) {
	return s.funcGetFollowers(ctx, username)
}
func (s mockService) GetFollowings(ctx context.Context, username string) (*dto.Followings, error) {
	return s.funcGetFollowings(ctx, username)
}
func (s mockService) GetTimeline(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
	return s.funcGetTimeline(ctx, userCtx)
}
func (s mockService) CreateGeneralNotifications(ctx context.Context, req *dto.AdminNotificationRequest) error {
	return s.funcCreateGeneralNotifications(ctx, req)
}
func (s mockService) GetNotificationsByUsername(ctx context.Context, username string) (*dto.NotificationsResponse, error) {
	return s.funcGetNotificationsByUsername(ctx, username)
}
