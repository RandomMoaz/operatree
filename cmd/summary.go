package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
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
	fmt.Printf("Default project: %s (%s)\n", cfg.Default.Name, cfg.Default.AbsPath)

	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}
	p.Summary()
}
