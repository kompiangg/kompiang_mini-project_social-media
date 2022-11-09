package impl

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	cloudinary "github.com/kompiang_mini-project_social-media/pkg/utils/cloudinaryutils"
)

func (s service) CreatePost(ctx context.Context, req *dto.CreatePostRequest, username string, imageFileName *string, videoFileName *string) (*dto.Post, error) {
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
		req.PictureURL = imageURL
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
		req.VideoURL = videoURL
	}

	res, err := s.repository.CreatePost(ctx, req, username, id.String())
	if err != nil {
		log.Println("[CreatePost Service Error] While calling create post repository:", err.Error())
		return nil, err
	}
	return res, nil
}

func (s service) GetAllPostByUsername(ctx context.Context, username string) (*dto.Posts, error) {
	res, err := s.repository.GetAllPostByUsername(ctx, username)
	if err != nil {
		log.Println("[GetAllPostByUsername Service Error] While calling get all post by username repository:", err.Error())
		return nil, err
	}
	return res, nil
}

func (s service) DeletePost(ctx context.Context, req *dto.DeletePostRequest) error {
	err := s.repository.DeletePost(ctx, req)
	if err != nil {
		log.Println("[DeletePost Service Error] While calling create post repository:", err.Error())
		return err
	}
	return nil
}

func (s service) GetPostByID(ctx context.Context, id string) (*dto.Post, error) {
	res, err := s.repository.GetPostByID(ctx, id)
	if err != nil {
		log.Println("[GetPostByID Service Error] While calling get post by id repository:", err.Error())
		return nil, err
	}
	return res, nil
}
