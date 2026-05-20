package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var plain bool

func init() {
	descCmd.Flags().BoolVarP(&plain, "plain", "p", false, "output raw YAML for piping")
	rootCmd.AddCommand(descCmd)
}

var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "Describes a project",
	Long:  "Describes a project and prints its metadata",
	Args:  cobra.NoArgs,
	Run:   describe,
}

func describe(cmd *cobra.Command, args []string) {
	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}

	p.Describe(plain)
}
