//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"

	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/google/wire"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

func NewSephirahService(*conf.Sephirah_Data, *mapper.LibrarianMapperServiceClient, *searcher.LibrarianSearcherServiceClient, *porter.LibrarianPorterServiceClient) (pb.LibrarianSephirahServiceServer, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
