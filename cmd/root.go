package cmd

import (
	"log"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/internal/config"
	"github.com/spf13/cobra"
)

var (
	// cfgFile store the configuration file name
	//cfgFile                 string
	//verbose, prettyPrintLog bool
	rootCmd = &cobra.Command{
		Use:   "go-app",
		Short: "Backend Server",
		Long:  `Backend Server`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.Println("Loading configurations")
	config.Init()
	log.Println("Configurations loaded successfully!")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
