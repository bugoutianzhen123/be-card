//go:build wireinject

package main

import (
	"github.com/asynccnu/be-card/grpc"
	"github.com/asynccnu/be-card/ioc"
	"github.com/asynccnu/be-card/pkg/grpcx"
	"github.com/asynccnu/be-card/repository"
	"github.com/asynccnu/be-card/repository/cache"
	"github.com/asynccnu/be-card/repository/dao"
	"github.com/asynccnu/be-card/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewCardGrpcService,
		service.NewCardService,
		repository.NewCardRepository,
		dao.NewCardDao,
		cache.NewCardRedisCache,
		// 第三方
		ioc.InitEtcdClient,
		ioc.InitRedis,
		ioc.InitDB,
		ioc.InitLogger,
	)
	return grpcx.Server(nil)
}
