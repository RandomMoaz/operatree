package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var showTracked bool
var tmpltName string

func init() {
	trackCmd.Flags().BoolVar(&showTracked, "show", false, "show tracked projects")
	trackCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	trackCmd.MarkFlagsOneRequired("dest", "show")
	trackCmd.PreRun = resolveProjectDir
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
	// Load config
	c, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load config file. Use operatree init to initialize config"))
	}

	if showTracked {
		c.ListProjects()
		return
	}

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
