package deteccted

import (
	"fmt"
	"github.com/detecc/deteccted-v2/internal/detecc"
	"github.com/detecc/deteccted-v2/internal/pkg/configuration"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

const (
	debugFlag = "debug"
)

var (
	configurationFilePathFlag string
	isDebug                   = false

	rootCmd = &cobra.Command{
		Use:   "deteccted",
		Short: "Deteccted V2 is a client for Detecctor V2.",
		Long:  ``,
		Run:   run,
	}
)

func init() {
	var (
		workingDirectory, _   = os.Getwd()
		defaultConfigFileName = fmt.Sprintf("%s/client.%s", workingDirectory, "yaml")
	)

	// Set flags
	rootCmd.PersistentFlags().StringVarP(&configurationFilePathFlag, "config", "c", defaultConfigFileName, "config file path")
	rootCmd.PersistentFlags().BoolVarP(&isDebug, debugFlag, "d", false, "debug mode")
}

func run(cmd *cobra.Command, args []string) {
	// Get configuration
	cfg := configuration.GetClientConfiguration(configurationFilePathFlag)

	// Run the app based on the configuration
	detecc.Run(isDebug, cfg)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Cannot run the client")
	}
}
