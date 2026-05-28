package project

import (
	"path"

	"github.com/hanymamdouh82/operatree/pkg/module"
)

func hydratePath(projectBaseDir string, p *Project) {
	p.absDir = projectBaseDir
	for i, m := range p.Modules {
		p.Modules[i].AbsPath = path.Join(projectBaseDir, m.Name)
		hydrateModule(&p.Modules[i])
	}
}

func hydrateModule(m *module.Module) {

	for i, s := range m.Subjects {
		m.Subjects[i].DirName = path.Join(m.AbsPath, s.Name)
	}

	for i, sm := range m.Modules {
		m.Modules[i].AbsPath = path.Join(m.AbsPath, sm.Name)
		hydrateModule(&m.Modules[i])
	}
}
