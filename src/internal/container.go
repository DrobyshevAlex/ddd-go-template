package internal

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/api"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/config"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/middleware"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/local"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/user"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	processError(container.Provide(config.New))
	processError(container.Provide(New))

	// Data
	processError(container.Provide(user.New))
	processError(container.Provide(local.New))
	processError(container.Provide(repo.New))

	// GRPC Service
	processError(container.Provide(func() *grpc.Server {
		return grpc.NewServer(
			grpc.ChainUnaryInterceptor(middleware.ErrorInterceptor),
		)
	}))
	processError(container.Provide(api.New))

	return container
}

func processError(err error) {
	if err != nil {
		panic(errors.Wrap(err, "cannot provide container"))
	}
}
