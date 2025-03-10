package server

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libjwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Sephirah_Server, greeter pb.LibrarianSephirahServiceServer) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.GetLogger()),
			ratelimit.Server(),
			jwt.Server(
				libjwt.KeyFunc(""),
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				jwt.WithClaims(libjwt.NewClaims),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianSephirahServiceServer(srv, greeter)
	return srv
}
