package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/urfave/negroni"

	"github.com/sirupsen/logrus"

	"github.com/faizalpribadi/layout/internal/config"

	"github.com/faizalpribadi/layout/internal/domain"
)

type Server struct {
	user domain.UserHandler
	feed domain.FeedHandler
	cfg  config.Config
}

func NewServer(user domain.UserHandler, feed domain.FeedHandler, cfg config.Config) domain.Server {
	return &Server{user: user, feed: feed, cfg: cfg}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/feeds", s.feed.GetFeeds)
	mux.HandleFunc("/users", s.user.GetUsers)
	return mux
}

func (s *Server) gracefullShutdown(server *http.Server, logger *logrus.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

func (s *Server) Run() {
	c, err := s.cfg.ReadFileConfiguration("./configs")
	if err != nil {
		logrus.Error(err)
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	logger := logrus.New()

	n := negroni.Classic()
	n.UseHandler(s.Handler())
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.Http.Port),
		Handler: n,
	}

	go s.gracefullShutdown(server, logger, quit, done)
	logger.Println("Server is ready to handle requests at", c.Http.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", c.Http.Port, err)
	}

	<-done
	logger.Println("Server stopped")
}
