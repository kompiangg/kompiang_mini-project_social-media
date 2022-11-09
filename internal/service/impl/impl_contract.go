package impl

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/internal/repository"
)

type service struct {
	repository repository.Repository
	config     *config.Config
	cloudinary *cloudinary.Cloudinary
}

type ServiceParams struct {
	Repository repository.Repository
	Config     *config.Config
	Cloudinary *cloudinary.Cloudinary
}

func NewService(params ServiceParams) *service {
	return &service{
		repository: params.Repository,
		config:     params.Config,
		cloudinary: params.Cloudinary,
	}
}
