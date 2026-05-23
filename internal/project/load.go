package project

import (
	"os"
	"path"

	"github.com/hanymamdouh82/operatree/internal/module"
	"gopkg.in/yaml.v3"
)

// Loads a project by reading project metadata file and sets project structrue
// Path represents project root path
func Load(pth string) (Project, error) {

	b, err := os.ReadFile(path.Join(pth, METADATA_FILE))
	if err != nil {
		return Project{}, err
	}

	// unmarshal into loader struct, this is because Unit is an interface not a struct
	// then we can convert units as per build logic
	var p Project
	if err := yaml.Unmarshal(b, &p); err != nil {
		return Project{}, err
	}

	hydratePath(pth, &p)

	return p, err
}

func hydratePath(projectBaseDir string, p *Project) {
	p.BaseDir = projectBaseDir
	for i := range p.Modules {
		hydrateModule(projectBaseDir, &p.Modules[i])
	}
}

func hydrateModule(projectBaseDir string, m *module.Module) {

	// hydrate module abs path
	m.AbsPath = path.Join(projectBaseDir, m.Name)

	for i, s := range m.Subjects {
		m.Subjects[i].DirName = path.Join(projectBaseDir, m.Name, s.Name)
	}

	for i := range m.Modules {
		hydrateModule(projectBaseDir, &m.Modules[i])
	}
}
