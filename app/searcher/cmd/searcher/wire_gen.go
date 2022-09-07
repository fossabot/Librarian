// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/app/searcher/internal/server"
	"github.com/tuihub/librarian/app/searcher/internal/service"
	"github.com/tuihub/librarian/internal/conf"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(searcher_Server *conf.Searcher_Server, searcher_Data *conf.Searcher_Data) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(searcher_Data)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo)
	librarianSearcherServiceServer := service.NewLibrarianSearcherServiceService(greeterUsecase)
	grpcServer := server.NewGRPCServer(searcher_Server, librarianSearcherServiceServer)
	app := newApp(grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
