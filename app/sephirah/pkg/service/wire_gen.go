// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/protos/pkg/librarian/mapper/v1"
	v1_3 "github.com/tuihub/protos/pkg/librarian/porter/v1"
	v1_2 "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	v1_4 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

// Injectors from wire.go:

func NewSephirahService(sephirah_Data *conf.Sephirah_Data, librarianMapperServiceClient *v1.LibrarianMapperServiceClient, librarianSearcherServiceClient *v1_2.LibrarianSearcherServiceClient, librarianPorterServiceClient *v1_3.LibrarianPorterServiceClient) (v1_4.LibrarianSephirahServiceServer, func(), error) {
	client, cleanup, err := data.NewSqlClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(client)
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, librarianMapperServiceClient, librarianSearcherServiceClient, librarianPorterServiceClient)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(greeterUsecase)
	return librarianSephirahServiceServer, func() {
		cleanup()
	}, nil
}
