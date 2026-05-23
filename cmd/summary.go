package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	summaryCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	summaryCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summarizes project at a glance",
	Long:  "Prints a high-level summary of project subjects, counts and status",
	Args:  cobra.NoArgs,
	Run:   summary,
}

func summary(cmd *cobra.Command, args []string) {
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}
	p.Summary()
}
