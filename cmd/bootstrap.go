package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	cfg, _ := config.Load() // best effort, empty if no config yet
	defaultDir := "."
	if cfg.StandardDir != "" {
		defaultDir = cfg.StandardDir
	}

	bootstrapCmd.Flags().StringVarP(&baseDir, "base", "b", defaultDir, "project base directory")
	bootstrapCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show operation output")
	rootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap [project_name]",
	Short: "Bootstraps new project",
	Long:  `Bootstraps new porject in current working directory`,
	Args:  cobra.ExactArgs(1),
	Run:   bootstrap,
}

func bootstrap(cmd *cobra.Command, args []string) {
	pn := args[0]
	p, err := project.Bootstrap(pn, baseDir)
	if err != nil {
		log.Fatal(err)
	}

	// To-Do: add flag for verbose / silent
	if verbose {
		if err := p.Describe(false); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Project: %s (%s)\n", p.Name, p.ProjectDir)
}
