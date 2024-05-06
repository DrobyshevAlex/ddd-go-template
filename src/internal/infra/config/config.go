package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const ENV_PREFIX = ""

type Config struct {
	GrpcPort uint16 `envconfig:"GRPC_PORT" validate:"required"`
}

func New() (*Config, error) {
	_ = godotenv.Overload()

	c := &Config{}

	if err := envconfig.Process(ENV_PREFIX, c); err != nil {
		return nil, errors.Wrap(err, "cannot get env config")
	}

	if err := validator.New().Struct(c); err != nil {
		return nil, err
	}

	return c, nil
}
