package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	untrackCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)

	if err := untrackCmd.MarkFlagRequired("dest"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(untrackCmd)
}

var untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Untracks project",
	Long:  "Untracks current project from tracked projects",
	Args:  cobra.NoArgs,
	Run:   untrack,
}

func untrack(cmd *cobra.Command, args []string) {
	resolveProjectDir(cmd, args)

	if err := config.RemoveProject(actDir); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project untracked: %s\n", actDir)
}
