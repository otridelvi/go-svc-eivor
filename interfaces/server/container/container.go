package container

import "github.com/otridelvi/go-svc-eivor/internal/shared/config"

type Container struct {
	Config *config.Config
}

func Setup() *Container {
	cfg := config.NewConfig("./resources/config.json")

	return &Container{
		Config: cfg,
	}
}
