package dto

type FollowRequest struct {
	FollowRequestToUser string `json:"follow_request_to_user"`
}

type FollowResponse struct {
	FollowedUser string `json:"followed_user"`
}

type Followers []UserBrief
type Followings []UserBrief
