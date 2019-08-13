package cmd

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/faizalpribadi/layout/internal/container"
)

//HTTPServer struct
type HTTPServer struct{}

//NewHTTPApi initialize
func NewHTTPApi() *HTTPServer {
	return &HTTPServer{}
}

//Start function
func (s *HTTPServer) Start() {
	c, err := container.NewContainer(context.Background())
	if err != nil {
		logrus.Errorf("server error %s", err.Error())
	}
	c.Run()
}
