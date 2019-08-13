package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//Start command line
func Start() {
	var (
		rootCommand      = &cobra.Command{}
		serveHTTPCommand = &cobra.Command{
			Use:   "serve-http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				httpServer := NewHTTPApi()
				httpServer.Start()
			},
		}
	)

	rootCommand.AddCommand(serveHTTPCommand)
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
