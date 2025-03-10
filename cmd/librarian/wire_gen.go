// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/mapper/pkg/service"
	service3 "github.com/tuihub/librarian/app/porter/pkg/service"
	service2 "github.com/tuihub/librarian/app/searcher/pkg/service"
	service4 "github.com/tuihub/librarian/app/sephirah/pkg/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, mapper_Data *conf.Mapper_Data, searcher_Data *conf.Searcher_Data, porter_Data *conf.Porter_Data) (*kratos.App, func(), error) {
	librarianMapperServiceServer, cleanup, err := service.NewMapperService(mapper_Data)
	if err != nil {
		return nil, nil, err
	}
	librarianMapperServiceClient := inprocgrpc.NewInprocMapperChannel(librarianMapperServiceServer)
	librarianSearcherServiceServer, cleanup2, err := service2.NewSearcherService(searcher_Data)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceClient := inprocgrpc.NewInprocSearcherChannel(librarianSearcherServiceServer)
	librarianPorterServiceServer, cleanup3, err := service3.NewPorterService(porter_Data)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceClient := inprocgrpc.NewInprocPorterChannel(librarianPorterServiceServer)
	librarianSephirahServiceServer, cleanup4, err := service4.NewSephirahService(sephirah_Data, librarianMapperServiceClient, librarianSearcherServiceClient, librarianPorterServiceClient)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcServer := server.NewGRPCServer(sephirah_Server, librarianSephirahServiceServer)
	app := newApp(grpcServer)
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
