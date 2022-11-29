package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"tcmlabs.fr/hexagonal_architecture_golang/config"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/core/domain"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/primary_adapter/_http"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/secondary_adapter/repositories/user"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/cinema/secondary_adapter/repositories/user/mongo"
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

	serverCfg, err := _http.NewServerCfg(cfg.Http.Server.Address, uint16(cfg.Http.Server.Port))
	if err != nil {
		os.Exit(1)
	}

	userRepo, err := NewUserRepository(ctx, cfg.UserRepository.Implementation, cfg.UserRepository.Config)
	if err != nil {
		os.Exit(1)
	}

	userSvc := services.NewUserServices(userRepo)

	server := _http.NewServer(serverCfg, userSvc)

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

		log.Printf("mongo cinema repositories set : %+#v", userRepo)
		return userRepo, nil
	default:
		err := fmt.Errorf("unknown implementation request : %s", implRequest)
		log.Print(err)
		return nil, err
	}
}
