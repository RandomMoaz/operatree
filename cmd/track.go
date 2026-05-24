package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	trackCmd.Flags().StringVarP(&destDir, "dest", "d", "", dFlagHelp_baseDir)

	if err := trackCmd.MarkFlagRequired("dest"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track project",
	Long:  "Adds project to tracked projects",
	Args:  cobra.NoArgs,
	Run:   track,
}

func track(cmd *cobra.Command, args []string) {
	resolveProjectDir(cmd, args)

	// load project to confirm its state
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := config.AddProject(p.Name, p.ProjectDir(), p.Template); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project tracked: %s\n", actDir)
}
