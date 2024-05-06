package middleware

import (
	"context"
	"errors"

	domainErrors "github.com/DrobyshevAlex/ddd-go-template/internal/domain/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	res, err := handler(ctx, req)
	if err == nil {
		return res, nil
	}
	var (
		ve  domainErrors.ValidationError
		nf  domainErrors.NotFoundError
		aee domainErrors.AlreadyExistsError
		na  domainErrors.NotAllowedError
		ue  domainErrors.UnauthenticatedError
	)

	code := codes.Internal
	msg := err.Error()

	switch {
	case errors.As(err, &ve):
		code = codes.InvalidArgument
	case errors.As(err, &nf):
		code = codes.NotFound
	case errors.As(err, &aee):
		code = codes.AlreadyExists
	case errors.As(err, &na):
		code = codes.PermissionDenied
	case errors.As(err, &ue):
		code = codes.Unauthenticated
	default:
		msg = "Internal server error"
	}

	return nil, status.Error(code, msg)
}
