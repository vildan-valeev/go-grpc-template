// Package httpserver implements HTTP server.
package server

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog/log"
	"go-grpc-template/internal/config"
	"go-grpc-template/internal/domain/service"
	bff "go-grpc-template/internal/transport/grpc/bff"
	"go-grpc-template/pkg/errors"
	"go-grpc-template/pkg/logger/grpcadapter"
	pb "go-grpc-template/proto/generated"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	grpc   *grpc.Server
	config config.Config

	CRUDService pb.CRUDServer
}

// New returns a new instance of Server.
func New(cfg config.Config, services *service.Services) *Server {
	s := &Server{
		config: cfg,
	}
	var (
		payloadLoggingDecider = func(ctx context.Context, fullMethodName string, servingObject interface{}) logging.PayloadDecision {
			return logging.LogPayloadRequestAndResponse
		}
		optsRecovery = []recovery.Option{
			recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				log.Error().Msgf("panic triggered: %v", p)
				return bff.GRPCError(&errors.Error{Code: errors.Unknown, Message: "panic"})
			}),
		}
		optsLogger = []logging.Option{
			logging.WithLevels(logging.DefaultServerCodeToLevel),
		}
	)
	s.grpc = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			ZerologCtxUnaryServerInterceptor(log.Logger),
			RequestIDCtxUnaryServerInterceptor(),
			logging.UnaryServerInterceptor(grpcadapter.InterceptorLogger(log.Logger), optsLogger...),
			InjectRequestIDCtxUnaryServerInterceptor(),
			logging.PayloadUnaryServerInterceptor(grpcadapter.InterceptorLogger(log.Logger), payloadLoggingDecider, time.RFC3339),
			recovery.UnaryServerInterceptor(optsRecovery...),
		),
	)

	return s
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() error {

	// регистрация сервисов в сервере
	pb.RegisterCRUDServer(s.grpc, s.CRUDService)

	address := net.JoinHostPort(s.config.IP, s.config.GRPCPort)

	lis, err := net.Listen("tcp4", address)
	if err != nil {
		return err
	}

	go func() {
		log.Info().Msgf("Start GRPC on %s", address)

		if err := s.grpc.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to grpc serve")
		}
	}()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	s.grpc.GracefulStop()
	return nil
}
