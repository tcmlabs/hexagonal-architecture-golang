package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/config"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/core/services"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/primary_adapter/http"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/secondary_adapter/repository/user"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/secondary_adapter/repository/user/mongo"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("configuration set : %+#v", cfg)

	serverCfg, err := http.NewServerCfg(cfg.Http.Server.Address, uint16(cfg.Http.Server.Port))
	if err != nil {
		os.Exit(1)
	}

	userRepo, err := NewUserRepository(ctx, cfg.UserRepository.Implementation, cfg.UserRepository.Config)
	if err != nil {
		os.Exit(1)
	}

	userSvc := services.NewUserServices(userRepo)

	server := http.NewServer(serverCfg, userSvc)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func NewUserRepository(ctx context.Context, implRequest string, cfg string) (user.Repository, error) {
	switch implRequest {
	case "inmemory":
		//TODO implement me
		panic("implement me")
	case "mongo":
		mongoCfg, err := mongo.NewMongoCfg(cfg)
		if err != nil {
			return nil, err
		}

		userRepo, err := mongo.NewUserRepository(ctx, mongoCfg)
		if err != nil {
			return nil, err
		}

		log.Printf("mongo user repository set : %+#v", userRepo)
		return userRepo, nil
	default:
		err := fmt.Errorf("unknown implementation request : %s", implRequest)
		log.Print(err)
		return nil, err
	}
}
