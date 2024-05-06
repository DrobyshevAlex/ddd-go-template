package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/DrobyshevAlex/ddd-go-template/internal/gen/pb"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/api"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	cnf        *config.Config
	grpcServer *grpc.Server
	service    *api.Service
}

func New(cnf *config.Config, grpcServer *grpc.Server, service *api.Service) *Application {
	return &Application{
		cnf:        cnf,
		grpcServer: grpcServer,
		service:    service,
	}
}

func (a *Application) Run(ctx context.Context) error {
	reflection.Register(a.grpcServer)
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cnf.GrpcPort))

	if err != nil {
		return fmt.Errorf("Not init GRPC Listener %w", err)
	}

	pb.RegisterTaskServiceServer(a.grpcServer, a.service)

	go func() {
		fmt.Printf("Run grpc on port %d\n", a.cnf.GrpcPort)
		if err := a.grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("grpc server serve error")
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	a.grpcServer.GracefulStop()

	return nil
}
