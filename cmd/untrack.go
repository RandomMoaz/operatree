package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

func init() {
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

	if prjDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("cannot read current path")
		}
		prjDir = cwd
	}

	if err := config.RemoveProject(prjDir); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project untracked: %s\n", prjDir)
}
