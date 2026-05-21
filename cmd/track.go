package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

var showTracked bool
var tmpltName string

func init() {
	trackCmd.Flags().BoolVar(&showTracked, "show", false, "show tracked projects")
	trackCmd.Flags().StringVarP(&tmpltName, "template", "t", "", "template name")
	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track [project_name]",
	Short: "Track project",
	Long:  "Adds project to tracked projects",
	Args:  cobra.MaximumNArgs(1),
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

	var pn string

	if len(args) > 0 {
		pn = args[0]
	}

	if err := config.AddProject(pn, prjDir, tmpltName); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project tracked: %s\n", prjDir)
}
