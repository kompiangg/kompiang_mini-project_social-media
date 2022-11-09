package impl

import (
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
)

func (r repository) FollowOtherUser(ctx context.Context, req *dto.FollowRequest, userCtx *dto.UserContext) error {
	userRelations := models.UserRelation{
		FollowBy: userCtx.Username,
		FollowTo: req.FollowRequestToUser,
	}

	res := r.db.WithContext(ctx).Create(&userRelations)
	err := res.Error
	if err != nil {
		if errCode, ok := err.(*mysql.MySQLError); ok {
			if errCode.Number == 1062 {
				return errors.ErrAccountDuplicated
			}
		}
		log.Println("[FollowOtherUser Repository Error] While calling the query:", err.Error())
		return err
	}

	return nil
}

func (r repository) GetFollowers(ctx context.Context, username string) (*dto.Followers, error) {
	var userRelations []models.UserRelation

	res := r.db.WithContext(ctx).Where("follow_to = ?", username).Find(&userRelations)
	err := res.Error
	if err != nil {
		log.Println("[GetFollowers Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var followers dto.Followers
	for _, relation := range userRelations {
		follower := dto.UserBrief{
			Username: relation.FollowBy,
		}
		followers = append(followers, follower)
	}

	return &followers, nil
}

func (r repository) GetFollowings(ctx context.Context, username string) (*dto.Followings, error) {
	var userRelations []models.UserRelation

	res := r.db.WithContext(ctx).Where("follow_by = ?", username).Find(&userRelations)
	err := res.Error
	if err != nil {
		log.Println("[GetFollowings Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var followings dto.Followings
	for _, relation := range userRelations {
		follower := dto.UserBrief{
			Username: relation.FollowTo,
		}
		followings = append(followings, follower)
	}

	return &followings, nil
}
