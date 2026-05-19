package cmd

import (
	"fmt"
	"log"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"github.com/spf13/cobra"
)

var entityName string

func init() {
	newCmd.Flags().StringVarP(&prjDir, "dest", "d", "/mnt/extra/onfly/testprj", "project directory")
	newCmd.Flags().StringVarP(&entityName, "name", "n", "", "entity name")
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:       "new [event | task]",
	Short:     "creates new entity",
	Long:      "creates new entity within project",
	ValidArgs: []cobra.Completion{"event", "task"},
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
	Run:       newUnitEntity,
}

func newUnitEntity(cmd *cobra.Command, args []string) {
	a := args[0]
	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}

	switch a {
	case "event":
		if err := newEvent(&p); err != nil {
			log.Fatal(err)
		}

	case "task":
		fmt.Println("will create task")
	default:
		fmt.Println("i shouldn't be here")
	}
}

func newEvent(p *project.Project) error {

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == "events"
	})
	m := p.Modules[i]

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(subject.SubjectEvent, m.AbsPath)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	return nil
}
