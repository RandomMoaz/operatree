package cmd

import (
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var (
	rootDir string
)

func init() {
	bootstrapCmd.Flags().StringVarP(&rootDir, "dest", "d", "/mnt/extra/onfly", "project root directory")
	rootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap [project_name]",
	Short: "Bootstraps new project",
	Long:  `Bootstraps new porject in current working directory`,
	Args:  cobra.ExactArgs(1),
	Run:   bootstrap,
}

func bootstrap(cmd *cobra.Command, args []string) {
	pn := args[0]
	project.Bootstrap(rootDir, pn)
}
