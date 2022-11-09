package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/utils/passwordutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/tokenutils"
)

func (s service) LoginService(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserSuccessLoginResponse, error) {
	hashedPassword := passwordutils.HashPassword(req.Password)
	req.Password = hashedPassword

	res, err := s.repository.AccountLogin(ctx, req)
	if err != nil {
		log.Println("[AccountLogin Service] Error while calling the repository:", err.Error())
		return nil, err
	}

	refreshToken := tokenutils.NewRefreshToken(&tokenutils.RefreshTokenParams{
		Username: res.Username,
	})

	accessToken, err := tokenutils.NewAccessToken(s.config.JWTSecretKey, &dto.UserContext{
		Username:    res.Username,
		Email:       res.Email,
		DisplayName: res.DisplayName,
	})

	if err != nil {
		log.Println("[AccountLogin Service] Error while creating JWT access token:", err.Error())
		return nil, err
	}

	err = s.repository.InsertRefreshToken(ctx, refreshToken)
	if err != nil {
		log.Println("[AccountLogin Service] Error while inserting refresh token:", err.Error())
		return nil, err
	}

	return &dto.UserSuccessLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
	}, nil
}
