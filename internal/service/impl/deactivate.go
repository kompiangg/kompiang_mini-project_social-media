package impl

import (
	"context"
	"log"
)

func (s service) DeactivateUser(ctx context.Context, username string) error {
	err := s.repository.DeactivateUser(ctx, username)
	if err != nil {
		log.Println("[DeactivateUser Service] Error while calling the edit user repository:", err.Error())
		return err
	}
	return nil
}
