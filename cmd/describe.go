package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	descCmd.Flags().StringVarP(&prjDir, "dest", "d", "/mnt/extra/onfly/testprj", "project directory")
	rootCmd.AddCommand(descCmd)
}

var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "describes a project",
	Long:  "describes a project and prints its metadata",
	Args:  cobra.NoArgs,
	Run:   describe,
}

func describe(cmd *cobra.Command, args []string) {
	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}

	p.Describe()
}
