package main

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/client"
	"SomeRestApi/internal/config"
	handler "SomeRestApi/internal/handlers/hendler_REST"
	"SomeRestApi/internal/handlers/hendler_gRPC"
	"SomeRestApi/internal/handlers/hendler_gRPC/pb"
	"SomeRestApi/internal/repository"
	"SomeRestApi/internal/repository/postgres"
	"SomeRestApi/internal/service"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error initializating etc %s", err.Error())
	}

	var repos *repository.Repository
	if cfg.SaveType == "postgres" {
		db, err := postgres.InitConnect(cfg.Postgres)
		if err != nil {
			log.Fatalf("error initializing datatbase: %s", err.Error())
		}
		repos = repository.NewRepositoryPostgres(db)
	} else {
		repos = repository.NewRepositoryInMemory()
	}

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// REST SERVER
	go func() {
		srv := new(someRestApi.Server)
		if err := srv.Run(cfg.Server.PortREST, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	// gRPC SERVER
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterContactServiceServer(grpcServer, hendler_gRPC.NewContactHendler(services))
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.Server.PortGRPC))
		log.Printf(fmt.Sprintf("localhost:%s", cfg.Server.PortGRPC))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer.Serve(lis)
	}()

	// RabbitMQ
	ampqClient := client.NewClient(services)
	msg, err := ampqClient.RunClient(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
