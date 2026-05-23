package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var plain bool

func init() {
	descCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	descCmd.Flags().BoolVarP(&plain, "plain", "p", false, "output raw YAML for piping")
	descCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(descCmd)
}

var descCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describes a project",
	Long:  "Describes a project and prints its metadata",
	Args:  cobra.NoArgs,
	Run:   describe,
}

func describe(cmd *cobra.Command, args []string) {

	// fmt.Printf("destDir: %s\n", destDir)
	// fmt.Printf("actDir: %s\n", actDir)

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	p.Describe(plain)
}
