package impl

import (
	"context"
	"log"
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/tokenutils"
)

func (s service) RefreshTokenService(ctx context.Context, refreshToken *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	checkRefreshToken, err := s.repository.GetRefreshToken(ctx, refreshToken.RefreshToken)
	if err != nil {
		log.Println("[RefreshAccessToken Service] Error while calling the repository:", err.Error())
		return nil, err
	}

	if checkRefreshToken.ExpirationDate.Before(time.Now()) {
		log.Println("[RefreshAccessToken Service] The refresh token expired:", err.Error())
		return nil, errors.ErrUnauthorized
	}

	accCtx, err := s.repository.GetUserByUsername(ctx, checkRefreshToken.Username)
	if err != nil {
		log.Println("[RefreshAccessToken Service] Error while calling the find account by id repository:", err.Error())
		return nil, err
	}

	accessToken, err := tokenutils.NewAccessToken(s.config.JWTSecretKey, accCtx)
	if err != nil {
		log.Println("[RefreshAccessToken Service] Error while creating JWT access token:", err.Error())
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}
