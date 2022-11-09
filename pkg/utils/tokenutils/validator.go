package tokenutils

import (
	"encoding/base64"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
)

func ValidateAccessToken(conf *config.JWTConfig, accessToken string) (*dto.UserContext, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnauthorized
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.ErrUnauthorized
		}

		return []byte(conf.JWTSecretKey), nil
	})

	if err != nil {
		return nil, errors.ErrUnauthorized
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.ErrUnauthorized
	}
	user := dto.UserContext{
		Username:    mapClaims["username"].(string),
		Email:       mapClaims["email"].(string),
		DisplayName: mapClaims["display_name"].(string),
	}

	return &user, nil
}

func ValidateBasicToken(conf *config.Admin, basicToken string) error {
	decodedByteBasicToken, err := base64.StdEncoding.DecodeString(basicToken)
	if err != nil {
		return err
	}

	decodedStringBasicToken := string(decodedByteBasicToken)
	splitted := strings.Split(decodedStringBasicToken, ":")
	if len(splitted) != 2 {
		return errors.ErrUnauthorized
	}

	username, password := splitted[0], splitted[1]
	if username != conf.Username || password != conf.Password {
		return errors.ErrUnauthorized
	}

	return nil
}
