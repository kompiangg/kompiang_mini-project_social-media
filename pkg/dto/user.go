package dto

import (
	"time"
)

// Register

type UserRegisterRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
}

// Login

type UserSuccessLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username    string
	Email       string
	DisplayName string
}

// Profile

type EditProfileRequest struct {
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	DisplayName       string  `json:"display_name"`
	ProfilePictureURL *string `json:"profile_picture"`
	Bio               string  `json:"bio"`
	BirthDate         string  `json:"birth_date"`
}

type EditProfileResponse struct {
	Username          string     `json:"username"`
	Email             string     `json:"email"`
	DisplayName       string     `json:"display_name"`
	ProfilePictureURL string     `json:"profile_picture"`
	Bio               string     `json:"bio"`
	BirthDate         *time.Time `json:"birth_date"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type UserMyProfileResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	CountFollowers    int       `json:"count_followers"`
	CountFollowings   int       `json:"count_followings"`
	DisplayName       string    `json:"display_name"`
	ProfilePictureURL string    `json:"profile_picture"`
	Bio               string    `json:"bio"`
	IsVerifiedEmail   bool      `json:"is_verified_email"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UserOtherProfileResponse struct {
	Username          string    `json:"username"`
	CountFollowers    int       `json:"count_followers"`
	CountFollowings   int       `json:"count_followings"`
	DisplayName       string    `json:"display_name"`
	ProfilePictureURL string    `json:"profile_picture"`
	Bio               string    `json:"bio"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UserBrief struct {
	Username string `json:"username"`
	// DisplayName       string `json:"display_name"`
	// ProfilePictureURL string `json:"profile_picture_url"`
}
