package impl

import (
	"context"
	goerrors "errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"gorm.io/gorm"
)

func (r repository) AccountLogin(ctx context.Context, req *dto.UserLoginRequest) (*dto.UserLogin, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := r.db.WithContext(ctx).
		Preload("UserDetail").
		Preload("UserSetting").
		Where(models.UserSetting{IsDeactivate: false}).
		First(&user).Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUnauthorized
		}
		log.Println("[AccountLogin Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	return &dto.UserLogin{
		Username:    user.Username,
		Email:       user.Email,
		DisplayName: user.UserDetail.DisplayName,
	}, nil
}

func (r repository) InsertRefreshToken(ctx context.Context, refreshToken *dto.UserRefreshToken) error {
	refreshTokenModel := models.RefreshToken{
		Username:       refreshToken.Username,
		Token:          refreshToken.Token,
		ExpirationDate: refreshToken.ExpirationDate,
	}

	res := r.db.WithContext(ctx).Create(&refreshTokenModel)
	err := res.Error
	if err != nil {
		log.Println("[AccountLogin Repository Error] While calling the query:", err.Error())
		return err
	}

	return nil
}

func (r repository) RegisterUser(ctx context.Context, req *dto.UserRegisterRequest) error {
	userModel := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		UserDetail: models.UserDetail{
			Username:          req.Username,
			DisplayName:       req.DisplayName,
			ProfilePictureURL: "",
			Bio:               req.Bio,
			BirthDate:         nil,
		},
		UserSetting: models.UserSetting{
			Username:        req.Username,
			IsVerifiedEmail: false,
			IsDeactivate:    false,
		},
	}

	res := r.db.WithContext(ctx).Create(&userModel)
	err := res.Error

	if errCode, ok := err.(*mysql.MySQLError); ok {
		if errCode.Number == 1062 {
			return errors.ErrAccountDuplicated
		}
	}

	if err != nil {
		log.Println("[CreateUser Repository]", err.Error())
		return err
	}

	return nil
}

func (r repository) GetRefreshToken(ctx context.Context, refreshToken string) (*dto.UserRefreshToken, error) {
	var refreshTokenModel models.RefreshToken

	res := r.db.WithContext(ctx).First(&refreshTokenModel, "token = ?", refreshToken)
	err := res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUnauthorized
		}
		return nil, err
	}

	return &dto.UserRefreshToken{
		Username:       refreshTokenModel.Username,
		Token:          refreshToken,
		ExpirationDate: refreshTokenModel.ExpirationDate,
		IsInvalidated:  false,
	}, nil
}

func (r repository) GetUserByUsername(ctx context.Context, username string) (*dto.UserContext, error) {
	user := models.User{
		Username: username,
	}

	res := r.db.WithContext(ctx).Preload("UserDetail").First(&user)
	err := res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[CreateUser Repository]", err.Error())
			return nil, errors.ErrUnauthorized
		}
		log.Println("[CreateUser Repository]", err.Error())
		return nil, err
	}

	return &dto.UserContext{
		Username:    username,
		Email:       user.Email,
		DisplayName: user.UserDetail.DisplayName,
	}, nil
}
