package project

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
// `bpth` is the abs base dir for the project, without project name included into it
func tmpltConsulting(name string, bpth string) Project {

	ppth := path.Join(bpth, name)

	p := Project{
		Name:    name,
		BaseDir: bpth,
		Modules: []module.Module{
			module.FactoryAdmin(ppth, "00"),
			module.FactoryEvents(ppth, "01"),
			module.FactoryLegal(ppth, "02"),
			module.FactoryResearch(ppth, "03"),
			module.FactoryDeliverables(ppth, "04"),
			module.FactoryArchive(ppth, "99"),
		},
	}

	return p
}
