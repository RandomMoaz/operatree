package project

import (
	"fmt"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/hanymamdouh82/operatree/pkg/module"
)

func Factory(tmplt template.OTTemplate, pname string) (Project, error) {
	if pname == "" {
		return Project{}, fmt.Errorf("error: project name is empty")
	}

	if len(tmplt.Modules) == 0 {
		return Project{}, fmt.Errorf("error: project doesn't include any modules")
	}

	p := Project{
		Name:     pname,
		Template: tmplt.Name,
	}

	// Range by index to pass the pointer of the template module securely
	for i := range tmplt.Modules {
		m, err := parseModule(&tmplt.Modules[i])
		if err != nil {
			return Project{}, err
		}

		p.Modules = append(p.Modules, m)
	}

	return p, nil
}

func parseModule(m *template.ModuleTemplate) (module.Module, error) {

	var moduleName string

	mt := strings.ToUpper(m.Type)
	mnu := strings.ToUpper(m.Name)
	mprfx, ok := module.ModuleDirPrefixMap[module.ModuleType(mt)]
	if !ok {
		return module.Module{}, fmt.Errorf("undefined module name prefix")
	}

	moduleName = fmt.Sprintf("%s_%s", mprfx, mnu)

	pm := module.Module{
		Type:    module.ModuleType(mt),
		Name:    moduleName,
		SubDirs: m.SubDirs,
		Modules: make([]module.Module, 0, len(m.Modules)),
	}

	// Recursively parse child modules and append to parent
	for i := range m.Modules {
		sm, err := parseModule(&m.Modules[i])
		if err != nil {
			return module.Module{}, err
		}
		pm.Modules = append(pm.Modules, sm)
	}

	return pm, nil
}
