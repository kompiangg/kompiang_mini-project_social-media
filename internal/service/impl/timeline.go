package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

func (s service) GetTimeline(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
	res, err := s.repository.GetTimeline(ctx, userCtx)
	if err != nil {
		log.Println("[GetTimeline Service Error] While calling get timeline repository:", err.Error())
		return nil, err
	}
	return res, nil
}
