package grpc

import (
	"github.com/owncloud/ocis-hello/pkg/proto/v0"
	"github.com/owncloud/ocis/ocis-pkg/service/grpc"
)

// Server initializes a new go-micro service ready to run
func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)
	handler := options.Handler

	service := grpc.NewService(
		grpc.Name(options.Config.Server.Name),
		grpc.Context(options.Context),
		grpc.Address(options.Config.GRPC.Addr),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Logger(options.Logger),
		grpc.Flags(options.Flags...),
		grpc.Version(options.Config.Server.Version),
	)

	if err := proto.RegisterHelloHandler(service.Server(), handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register service handler")
	}

	service.Init()
	return service
}
