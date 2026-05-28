package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize OperaTree configuration",
	Long:  "Creates initial OperaTree config file in OS config directory",
	Args:  cobra.NoArgs,
	Run:   initConfig,
}

func initConfig(cmd *cobra.Command, args []string) {
	if err := config.InitializeConfig(); err != nil {
		log.Fatal(err)
	}
}
