package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

func (s service) CreateGeneralNotifications(ctx context.Context, req *dto.AdminNotificationRequest) error {
	err := s.repository.CreateGeneralNotifications(ctx, req)
	if err != nil {
		log.Println("[CreateGeneralNotifications Service] Error while calling the repository:", err.Error())
		return err
	}
	return nil
}

func (s service) GetNotificationsByUsername(ctx context.Context, username string) (*dto.NotificationsResponse, error) {
	res, err := s.repository.GetNotificationsByUsername(ctx, username)
	if err != nil {
		log.Println("[GetNotificationsByUsername Service] Error while calling the repository:", err.Error())
		return nil, err
	}
	return res, nil
}
