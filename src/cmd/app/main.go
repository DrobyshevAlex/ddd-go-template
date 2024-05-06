package main

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal"
	"github.com/pkg/errors"
)

func main() {
	container := internal.BuildContainer()
	err := container.Invoke(func(application *internal.Application) {
		ctx := context.Background()
		err := application.Run(ctx)
		if err != nil {
			panic(errors.Wrap(err, "cannot run application"))
		}
	})

	if err != nil {
		panic(errors.Wrap(err, "application breakpoint"))
	}
}
