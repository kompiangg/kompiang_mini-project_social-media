package web

import (
	"log"

	"github.com/kompiang_mini-project_social-media/cmd/web/router"
	"github.com/kompiang_mini-project_social-media/config"
	repositorypkg "github.com/kompiang_mini-project_social-media/internal/repository/impl"
	servicepkg "github.com/kompiang_mini-project_social-media/internal/service/impl"
	"github.com/kompiang_mini-project_social-media/pkg/utils/websocketutils"
	"github.com/labstack/echo/v4"
)

type WebServiceParams struct {
	Config *config.Config
}

func StartWebService(params WebServiceParams) {
	db := config.GetDatabaseConn(params.Config.Database)
	defer func() {
		err := config.CloseDatabaseConnection()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("[INFO] Database connection closed gracefully")
	}()

	cloudinary, err := config.GetCloudinaryConn(&params.Config.Cloudinary)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	websocketPool := websocketutils.NewPool()
	go func() {
		websocketPool.Start()
	}()

	repository := repositorypkg.NewRepository(repositorypkg.RepositoryParams{
		DB: db,
	})

	service := servicepkg.NewService(servicepkg.ServiceParams{
		Repository: repository,
		Config:     params.Config,
		Cloudinary: cloudinary,
	})

	router.InitRoute(router.RouteParams{
		E:       e,
		Service: service,
		Config:  params.Config,
		Pool:    websocketPool,
	})

	err = config.StartServer(config.Server{
		E:    e,
		Port: params.Config.Server.Port,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
