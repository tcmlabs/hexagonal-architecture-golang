package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/netip"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/core/services"
)

const (
	ReadTimeout  = 15 * time.Second
	WriteTimeout = 15 * time.Second
)

type ServerCfg struct {
	addressPort netip.AddrPort
}

func NewServerCfg(address string, port uint16) (*ServerCfg, error) {
	addrValidated, err := netip.ParseAddr(address)
	if err != nil {
		err := fmt.Errorf("failed to parse address : %s, err : %w", address, err)
		log.Print(err)
		return nil, err
	}

	return &ServerCfg{
		addressPort: netip.AddrPortFrom(addrValidated, port),
	}, nil
}

func (s ServerCfg) Address() string {
	return s.addressPort.String()
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *ServerCfg, userService services.User) Server {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUser(userService)).Methods(http.MethodGet)
	router.HandleFunc("/users", createUser(userService)).Methods(http.MethodPost)

	httpServer := &http.Server{
		Addr:         cfg.addressPort.String(),
		Handler:      router,
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
	}

	return Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	log.Printf("server starting on %s", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		err := fmt.Errorf("failed to start server: %w", err)
		log.Print(err)
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	return err
}
