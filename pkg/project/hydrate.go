package project

import (
	"path/filepath"

	"github.com/hanymamdouh82/operatree/pkg/module"
)

func hydratePath(projectBaseDir string, p *Project) {
	p.absDir = filepath.Clean(projectBaseDir)
	for i, m := range p.Modules {
		p.Modules[i].AbsPath = filepath.Join(projectBaseDir, m.Name)
		hydrateModule(&p.Modules[i])
	}
}

func hydrateModule(m *module.Module) {

	for i, s := range m.Subjects {
		m.Subjects[i].DirName = filepath.Join(filepath.Clean(m.AbsPath), s.Name)
	}

	for i, sm := range m.Modules {
		m.Modules[i].AbsPath = filepath.Join(filepath.Clean(m.AbsPath), sm.Name)
		hydrateModule(&m.Modules[i])
	}
}
