package cmd

import (
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var (
	pDir string
)

func init() {
	loadCmd.Flags().StringVarP(&pDir, "dest", "d", "/mnt/extra/onfly/testprj", "project directory")
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "loads a project",
	Long:  "loads a project and prints its metadata",
	Args:  cobra.NoArgs,
	Run:   load,
}

func load(cmd *cobra.Command, args []string) {
	project.Load(pDir)
}
