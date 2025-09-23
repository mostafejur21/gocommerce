package rest

import (
	"fmt"
	"go_ecommerce/config"
	"go_ecommerce/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf *config.Config
}

func NewServer(cnf *config.Config) *Server {
	return &Server{cnf: cnf}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Starting server on port:", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
