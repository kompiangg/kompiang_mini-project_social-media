package main

import (
	"github.com/kompiang_mini-project_social-media/cmd/web"
	"github.com/kompiang_mini-project_social-media/config"
)

func main() {
	config := config.GetConfig()
	web.StartWebService(web.WebServiceParams{
		Config: config,
	})
}
