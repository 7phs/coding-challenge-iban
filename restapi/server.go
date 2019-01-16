package restapi

import (
	"context"
	"github.com/7phs/coding-challenge-iban/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	readTimeout     = 5 * time.Second
	writeTimeout    = 10 * time.Second
	shutdownTimeout = 5 * time.Second
)

type Server struct {
	http.Server

	prefix string
}

func NewServer(conf *config.Config, router http.Handler) *Server {
	return &Server{
		Server: http.Server{
			Addr:         conf.Address(),
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},

		prefix: "http",
	}
}

func (o *Server) Run() *Server {
	go func() {
		log.Info(o.prefix+": start listening ", o.Addr)
		if err := o.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(o.prefix+": failed to start listening ", o.Addr, ": ", err)
		}
	}()

	return o
}

func (o *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	log.Info(o.prefix + ": shutdown")
	if err := o.Server.Shutdown(ctx); err != nil {
		log.Error(o.prefix + ": failed to shutdown HTTP server")
	}
}
