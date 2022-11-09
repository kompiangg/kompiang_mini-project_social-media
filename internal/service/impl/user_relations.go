package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

func (s service) FollowOtherUser(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
	err := s.repository.FollowOtherUser(ctx, req, userCtx)
	if err != nil {
		log.Println("[FollowOtherUser Service Error] While calling get child comment repository:", err.Error())
		return err
	}
	return nil
}

func (s service) GetFollowers(ctx context.Context, username string) (*dto.Followers, error) {
	res, err := s.repository.GetFollowers(ctx, username)
	if err != nil {
		log.Println("[GetFollowers Service Error] While calling get child comment repository:", err.Error())
		return nil, err
	}
	return res, nil
}
func (s service) GetFollowings(ctx context.Context, username string) (*dto.Followings, error) {
	res, err := s.repository.GetFollowings(ctx, username)
	if err != nil {
		log.Println("[GetFollowers Service Error] While calling get child comment repository:", err.Error())
		return nil, err
	}
	return res, nil
}
