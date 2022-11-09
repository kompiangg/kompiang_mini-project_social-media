package tokenutils

import (
	"time"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/utils/randomutils"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenParams struct {
	Username string
}

func NewRefreshToken(params *RefreshTokenParams) *dto.UserRefreshToken {
	generatedToken := randomutils.GenerateNRandomString(128)
	return &dto.UserRefreshToken{
		Token:          generatedToken,
		Username:       params.Username,
		IsInvalidated:  false,
		ExpirationDate: time.Now().Add(time.Hour * 24 * 7),
	}
}

func NewAccessToken(jwtSecret string, acc *dto.UserContext) (string, error) {
	claims := dto.NewUserClaims(acc)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	accessToken, err := unsignedToken.SignedString([]byte(jwtSecret))
	return accessToken, err
}
