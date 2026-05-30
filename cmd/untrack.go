package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	untrackCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)

	rootCmd.AddCommand(untrackCmd)
}

var untrackCmd = &cobra.Command{
	Use:   "untrack [project_name]",
	Short: "Untracks project",
	Long:  "Untracks project from tracked projects",
	Args:  cobra.MaximumNArgs(1),
	Run:   untrack,
}

func untrack(cmd *cobra.Command, args []string) {

	// untrack using project name
	if len(args) > 0 {
		pn := args[0]
		if err := config.RemoveProjectByName(pn); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Project untracked: %s\n", pn)
		return
	}

	// Untrack using -d flag
	resolveProjectDir(cmd, args)
	if err := config.RemoveProject(actDir); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project untracked: %s\n", actDir)
}
