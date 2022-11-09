package impl

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	cloudinary "github.com/kompiang_mini-project_social-media/pkg/utils/cloudinaryutils"
)

func (s service) CreateComment(ctx context.Context, req *dto.CreateCommentRequest, username string, imageFileName *string, videoFileName *string) (*dto.CommentResponse, error) {
	id := uuid.New()

	if imageFileName != nil {
		imageURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
			Ctx:      ctx,
			Cld:      s.cloudinary,
			Filename: *imageFileName,
			Username: username,
		})
		if err != nil {
			return nil, err
		}
		req.PictureURL = *imageURL
	}

	if videoFileName != nil {
		videoURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
			Ctx:      ctx,
			Cld:      s.cloudinary,
			Filename: *videoFileName,
			Username: username,
		})
		if err != nil {
			return nil, err
		}
		req.VideoURL = *videoURL
	}

	res, err := s.repository.CreateComment(ctx, req, username, id.String())
	if err != nil {
		log.Println("[CreateComment Service Error] While calling create comment repository:", err.Error())
		return nil, err
	}
	return res, nil
}

func (s service) GetChildComment(ctx context.Context, parentID string) (*dto.CommentsResponse, error) {
	res, err := s.repository.GetChildComment(ctx, parentID)
	if err != nil {
		log.Println("[GetChildComment Service Error] While calling get child comment repository:", err.Error())
		return nil, err
	}
	return res, nil
}
