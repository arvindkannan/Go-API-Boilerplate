package cmd

import (
	// "go-api-boilerplate/config"
	"log"
	"go-api-boilerplate/server"
	"go-api-boilerplate/utils"
	"net"

	"github.com/spf13/cobra"
)

func InitializeApp(cmd *cobra.Command, _ []string) error {
	// load config
	// config := config.LoadConfig()

	// load logger
	utils.Logger()

	// Create a TCP listener on port 8000
	listener, _ := net.Listen("tcp", ":8000")
	// Server from server serve with the listener
	return server.Serve(listener)
}

func NewCLI() *cobra.Command {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:           "api-service",
		Short:         "API Service Runner",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cmd.UsageString())
		},
	}

	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	serveCmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"start"},
		Short:   "Start API Service",
		Args:    cobra.ExactArgs(0),
		RunE:    InitializeApp,
	}

	rootCmd.AddCommand(
		serveCmd,
	)

	return rootCmd

}
