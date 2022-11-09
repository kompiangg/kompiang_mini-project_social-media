package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/pkg/dto"
	cloudinary "github.com/kompiang_mini-project_social-media/pkg/utils/cloudinaryutils"
)

func (s service) EditUserService(ctx context.Context, userCtx *dto.UserContext, req *dto.EditProfileRequest, profilePictureFileName *string) (*dto.EditProfileResponse, error) {
	if profilePictureFileName != nil {
		imageURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
			Ctx:      ctx,
			Cld:      s.cloudinary,
			Filename: *profilePictureFileName,
			Username: userCtx.Username,
		})
		if err != nil {
			return nil, err
		}
		req.ProfilePictureURL = imageURL
	}

	res, err := s.repository.EditUser(ctx, req, userCtx)
	if err != nil {
		log.Println("[EditUserService Service] Error while calling the edit user repository:", err.Error())
		return nil, err
	}

	return res, nil
}

func (s service) GetOtherUserProfile(ctx context.Context, username string) (*dto.UserOtherProfileResponse, error) {
	profile, err := s.repository.GetOtherUserProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s service) GetMyProfile(ctx context.Context, username string) (*dto.UserMyProfileResponse, error) {
	profile, err := s.repository.GetMyProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
