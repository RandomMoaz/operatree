package cmd

import (
	"log"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:       "new [event | task | topic | objective]",
	Short:     "Creates new subject",
	Long:      "Creates new subject within project",
	ValidArgs: []cobra.Completion{"event", "task", "topic", "objective"},
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
		if err := newTask(&p); err != nil {
			log.Fatal(err)
		}

	case "topic":
		if err := newTopic(&p); err != nil {
			log.Fatal(err)
		}

	case "objective":
		if err := newObjective(&p); err != nil {
			log.Fatal(err)
		}

	default:
		return
	}
}

// Adds new event to Events module
func newEvent(p *project.Project) error {
	ss := project.ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleEvents
	})
	m := p.Modules[i]

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(subject.SubjectEvent, m.AbsPath, ss)
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

// Adds new event to Project Management / Tasks module
func newTask(p *project.Project) error {
	ss := project.ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleProjectManagement
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleTasks
	})

	m := pmm.Modules[j]

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(subject.SubjectTask, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

// Adds new topic to Research / Topics module
func newTopic(p *project.Project) error {
	ss := project.ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleResearch
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleTopics
	})

	m := pmm.Modules[j]

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(subject.SubjectTopic, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

// Adds new topic to Research / Objectives module
func newObjective(p *project.Project) error {
	ss := project.ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleResearch
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleObjectives
	})

	m := pmm.Modules[j]

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(subject.SubjectObjective, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	return nil
}
