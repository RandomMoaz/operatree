package project

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

// returns Dev project template
// `bpth` is the abs base dir for the project, without project name included into it
func dev(name string, bpth string) Project {

	ppth := path.Join(bpth, name)

	p := Project{
		Name:    name,
		BaseDir: bpth,
		Modules: []module.Module{
			{
				Type:     "admin",
				Name:     "00_ADMIN",
				AbsPath:  path.Join(ppth, "00_ADMIN"),
				Modules:  []module.Module{},
				Subjects: []subject.Subject{},
				SubDirs: []string{
					"contacts",
					"governance",
					"guidelines",
					"templates",
				},
			},
			{
				Type:     "events",
				Name:     "01_EVENTS",
				AbsPath:  path.Join(ppth, "01_EVENTS"),
				Modules:  []module.Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
			{
				Type:    "project_management",
				Name:    "02_PROJECT_MANAGEMENT",
				AbsPath: path.Join(ppth, "02_PROJECT_MANAGEMENT"),
				Modules: []module.Module{
					{
						Type:     "tasks",
						Name:     "01_TASKS",
						AbsPath:  path.Join(ppth, "02_PROJECT_MANAGEMENT", "01_TASKS"),
						Modules:  []module.Module{},
						Subjects: []subject.Subject{},
						SubDirs:  []string{},
					},
				},
				Subjects: []subject.Subject{},
				SubDirs: []string{
					"budgets",
					"communications",
					"planning",
					"reports",
					"risks",
				},
			},
		},
	}

	return p
}
