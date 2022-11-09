package impl

import (
	"context"
	goerrors "errors"
	"log"
	"sync"
	"time"

	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/dateutils"
	"gorm.io/gorm"
)

func (r repository) EditUser(ctx context.Context, req *dto.EditProfileRequest, userCtx *dto.UserContext) (*dto.EditProfileResponse, error) {
	var birthDate *time.Time
	var err error

	if req.BirthDate != "" {
		birthDate, err = dateutils.StringToDate(req.BirthDate)
		if err != nil {
			return nil, err
		}
	}

	user := models.User{
		Email: req.Email,
	}

	res := r.db.WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("username = ?", userCtx.Username).
		Model(&user).
		Updates(&user)

	err = res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		}
		log.Println("[EditUser Repository Error] While running the query:", err.Error())
		return nil, err
	}

	UserDetail := models.UserDetail{
		DisplayName:       req.DisplayName,
		ProfilePictureURL: *req.ProfilePictureURL,
		Bio:               req.Bio,
		BirthDate:         birthDate,
	}

	res = r.db.WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("username = ?", userCtx.Username).
		Model(&user.UserDetail).
		Updates(&UserDetail)

	err = res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		}
		log.Println("[EditUser Repository Error] While running the query:", err.Error())
		return nil, err
	}

	var newUser models.User
	res = r.db.WithContext(ctx).
		Preload("UserDetail").
		First(&newUser, "username = ?", userCtx.Username)

	err = res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return &dto.EditProfileResponse{
		Username:          newUser.Username,
		Email:             newUser.Email,
		DisplayName:       newUser.UserDetail.DisplayName,
		ProfilePictureURL: newUser.UserDetail.ProfilePictureURL,
		Bio:               newUser.UserDetail.Bio,
		BirthDate:         newUser.UserDetail.BirthDate,
		CreatedAt:         newUser.CreatedAt,
		UpdatedAt:         newUser.UpdatedAt,
	}, nil
}

func (r repository) DeactivateUser(ctx context.Context, username string) error {
	res := r.db.WithContext(ctx).
		Model(models.UserSetting{}).
		Where("username = ?", username).
		Update("is_deactivate", true)

	err := res.Error
	if err != nil {
		log.Println("[DeactivateUser Repository Error] While running the query:", err.Error())
		return err
	}
	if res.RowsAffected == 0 {
		return errors.ErrRecordNotFound
	}

	return nil
}

func (r repository) GetOtherUserProfile(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
	user := models.User{
		Username: username,
	}

	errChan := make(chan error)
	defer close(errChan)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		res := r.db.WithContext(ctx).
			Preload("UserSetting").
			Preload("UserDetail").
			Model(&user).
			First(&user).
			Where(models.UserSetting{IsDeactivate: false})
		err := res.Error
		if err != nil {
			if goerrors.Is(err, gorm.ErrRecordNotFound) {
				errChan <- errors.ErrRecordNotFound
				return
			}
			errChan <- err
			return
		}
		errChan <- nil
	}()

	var followersCount, followingsCount int64
	go func() {
		defer wg.Done()
		err := r.db.Model(&models.UserRelation{}).Where("follow_to = ?", username).Count(&followersCount).Error
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()
		err := r.db.Model(&models.UserRelation{}).Where("follow_by = ?", username).Count(&followingsCount).Error
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	for i := 0; i < 3; i++ {
		err := <-errChan
		if err != nil {
			wg.Wait()
			return nil, err
		}
	}

	return &dto.UserOtherProfileResponse{
		Username:          user.Username,
		DisplayName:       user.UserDetail.DisplayName,
		ProfilePictureURL: user.UserDetail.ProfilePictureURL,
		Bio:               user.UserDetail.Bio,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
		CountFollowers:    int(followersCount),
		CountFollowings:   int(followingsCount),
	}, nil
}

func (r repository) GetMyProfile(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
	user := models.User{
		Username: username,
	}

	errChan := make(chan error)
	defer close(errChan)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(errChan chan error) {
		defer wg.Done()
		res := r.db.WithContext(ctx).
			Preload("UserSetting").
			Preload("UserDetail").
			Model(&user).
			First(&user).
			Where(models.UserSetting{IsDeactivate: false})

		err := res.Error
		if err != nil {
			if goerrors.Is(err, gorm.ErrRecordNotFound) {
				errChan <- errors.ErrRecordNotFound
				return
			}
			errChan <- err
			return
		}
		errChan <- nil
	}(errChan)

	var followersCount, followingsCount int64
	go func(errChan chan error) {
		defer wg.Done()
		err := r.db.Model(&models.UserRelation{}).Where("follow_to = ?", username).Count(&followersCount).Error
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}(errChan)

	go func(errChan chan error) {
		defer wg.Done()
		err := r.db.Model(&models.UserRelation{}).Where("follow_by = ?", username).Count(&followingsCount).Error
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}(errChan)

	for i := 0; i < 3; i++ {
		err := <-errChan
		if err != nil {
			wg.Wait()
			return nil, err
		}
	}

	return &dto.UserMyProfileResponse{
		Username:          user.Username,
		Email:             user.Email,
		DisplayName:       user.UserDetail.DisplayName,
		ProfilePictureURL: user.UserDetail.ProfilePictureURL,
		Bio:               user.UserDetail.Bio,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
		IsVerifiedEmail:   user.UserSetting.IsVerifiedEmail,
		CountFollowers:    int(followersCount),
		CountFollowings:   int(followingsCount),
	}, nil
}
