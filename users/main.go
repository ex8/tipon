package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ex8/tipon/core/store"
	"github.com/ex8/tipon/users/pb"
	"github.com/ex8/tipon/users/svc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type cfg struct {
	Service serviceCfg
	Store   storeCfg
}

type serviceCfg struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type storeCfg struct {
	Client     string `yaml:"client"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

func main() {
	// config
	c := &cfg{}
	if file, err := os.Open("config.yml"); err == nil {
		defer file.Close()
		if err = yaml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	// logger
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer l.Sync()
	logger := l.Sugar()

	// listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Service.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// store
	s := store.New()

	// store ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// store connect
	opts := store.Opts{Client: c.Store.Client, Host: c.Store.Host, Port: c.Store.Port}
	if err = s.Connect(ctx, opts); err != nil {
		logger.Fatalf("failed to connect to store: %v", err)
	}

	// store disconnect
	defer func() {
		if err = s.Disconnect(ctx); err != nil {
			logger.Fatalf("failed to disconnect store: %v", err)
		}
	}()

	logger.Infof("starting %v service on port %v", c.Service.Name, c.Service.Port)

	// user grpc server
	server := grpc.NewServer()
	users := s.Client.Database(c.Store.Database).Collection(c.Store.Collection)
	pb.RegisterUserServiceServer(server, &svc.UserService{Users: users})
	if err := server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
