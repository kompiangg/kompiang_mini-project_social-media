package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/passwordutils"
)

func (s service) RegisterService(ctx context.Context, req *dto.UserRegisterRequest) error {
	hashedPassword := passwordutils.HashPassword(req.Password)
	req.Password = hashedPassword

	err := s.repository.RegisterUser(ctx, req)
	if err != nil {
		if err == errors.ErrAccountDuplicated {
			return errors.ErrAccountDuplicated
		}
		log.Println("[SERVICE ERROR] While calling register user repository")
		return err
	}

	return nil
}
