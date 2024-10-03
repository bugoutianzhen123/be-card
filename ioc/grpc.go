package ioc

import (
	lgrpc "github.com/asynccnu/be-card/grpc"
	"github.com/asynccnu/be-card/pkg/grpcx"
	"github.com/asynccnu/be-card/pkg/logger"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitGRPCxKratosServer(HelpServer *lgrpc.CardService, ecli *clientv3.Client, l logger.Logger) grpcx.Server {
	type Config struct {
		Name    string `yaml:"name"`
		Weight  int    `yaml:"weight"`
		Addr    string `yaml:"addr"`
		EtcdTTL int64  `yaml:"etcdTTL"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	server := kgrpc.NewServer(
		kgrpc.Address(cfg.Addr),
		kgrpc.Middleware(recovery.Recovery()),
	)

	HelpServer.Register(server)
	return &grpcx.KratosServer{
		Server:     server,
		Name:       cfg.Name,
		Weight:     cfg.Weight,
		EtcdTTL:    time.Second * time.Duration(cfg.EtcdTTL),
		EtcdClient: ecli,
		L:          l,
	}
}
