package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/spf13/cobra"
)

var templateName string

func init() {
	ts := make([]string, 0, len(template.Templates))
	for k := range template.Templates {
		ts = append(ts, k)
	}
	avts := strings.Join(ts, "|")
	fth := fmt.Sprintf("project template: %s", avts)

	bootstrapCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_baseDir)
	bootstrapCmd.Flags().StringVarP(&templateName, "template", "t", "", fth)
	bootstrapCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show operation output")

	if err := bootstrapCmd.MarkFlagRequired("template"); err != nil {
		log.Fatal(err)
	}

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
	// -d flag here is used to define base dir not project dir
	resolveBaseDir(cmd, args)

	pn := args[0]
	p, err := project.Bootstrap(pn, actDir, templateName)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		if err := p.Describe(false); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Project: %s (%s)\n", p.Name, p.ProjectDir())
}
